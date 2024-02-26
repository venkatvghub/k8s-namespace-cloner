package managers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"slices"
	"strconv"
	"strings"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/util/retry"
)

var (
	excludeSecretPrefixes    = []string{"sh.helm.release"}
	excludeConfigMapPrefixes = []string{"kube-root-ca.crt"}
)

type Deployment struct {
	Name      string
	Namespace string
	POD       string
	App       string
}

type Container map[string]string

type DeploymentDetail struct {
	Name       string      `json:"name"`
	Namespace  string      `json:"namespace"`
	POD        string      `json:"pod"`
	App        string      `json:"app"`
	Containers []Container `json:"containers"`
}

type DeploymentContainers struct {
	Deployments []DeploymentDetail `json:"deployments"`
}

func GetNS(clientset *kubernetes.Clientset) ([]map[string]string, *Error) {
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	namespaceNames := []map[string]string{}
	for _, namespace := range namespaces.Items {
		annotations := namespace.Annotations
		if annotations != nil {
			if _, ok := annotations[NS_CLONER_ANNOTATION]; ok {
				//log.Printf("Annotations:%v\n", annotations[NS_CLONER_ANNOTATION])
				if annotations[NS_CLONER_ANNOTATION] == "true" || annotations[NS_CLONER_ANNOTATION] == "True" {
					nsMap := make(map[string]string)
					nsMap["namespace"] = namespace.Name
					nsMap["Pod"] = namespace.Labels["POD"]
					nsMap["app"] = namespace.Labels["app"]
					namespaceNames = append(namespaceNames, nsMap)
				}
			}
		}
	}
	return namespaceNames, nil
}

func GetDeploymentForNS(clientset *kubernetes.Clientset, namespace string) ([]Deployment, *Error) {
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	deploymentObjects := []Deployment{}

	for _, deployment := range deployments.Items {
		d := Deployment{}
		replicas := *deployment.Spec.Replicas
		if replicas > 0 {
			d.Name = deployment.Name
			d.Namespace = deployment.Namespace
			d.POD = deployment.Labels["POD"]
			d.App = deployment.Labels["app"]
			deploymentObjects = append(deploymentObjects, d)
		}

	}
	if len(deploymentObjects) == 0 {
		return nil, &Error{
			Code:    http.StatusNotFound,
			Message: "No Deployments found eligible for display",
		}
	}
	return deploymentObjects, nil
}

func GetDeploymentYaml(clientset *kubernetes.Clientset, namespace string) (DeploymentContainers, *Error) {
	deployments, err := getDeploymentsForNS(clientset, namespace)
	deploymentContainers := DeploymentContainers{}
	if err != nil {
		return deploymentContainers, err
	}

	for _, deployment := range deployments.Items {
		// Initialize the inner map for each deployment
		// Only allow deployments that are cloned by this system using the annotations set
		errObj := validateDeploymentEliblity(clientset, &deployment)
		if errObj != nil {
			continue
		}
		d := DeploymentDetail{
			Name:      deployment.Name,
			Namespace: deployment.Namespace,
			POD:       deployment.Labels["POD"],
			App:       deployment.Labels["app"],
		}
		containers := []Container{}
		for _, container := range deployment.Spec.Template.Spec.Containers {
			c := make(Container)
			c[container.Name] = container.Image
			containers = append(containers, c)
		}
		d.Containers = containers
		deploymentContainers.Deployments = append(deploymentContainers.Deployments, d)
	}
	if len(deploymentContainers.Deployments) == 0 {
		return deploymentContainers, &Error{
			Code:    http.StatusNotFound,
			Message: "No Deployments found eligible for display",
		}
	}
	fmt.Printf("Deployments:%+v\n", deploymentContainers)
	return deploymentContainers, nil
}

func GetSecretYaml(clientset *kubernetes.Clientset, namespace string) (map[string]map[string]string, *Error) {
	secrets, err := getSecretsforNS(clientset, namespace)
	if err != nil {
		return nil, err
	}

	secretData := make(map[string]map[string]string)

	for _, secret := range secrets.Items {
		proceed := true
		errObj := validateSecretEliblity(clientset, &secret)
		if errObj != nil {
			continue
		}
		for _, name := range excludeSecretPrefixes {
			if strings.HasPrefix(secret.Name, name) {
				proceed = false
				continue
			}
		}

		if !proceed {
			continue
		}
		if slices.Contains(excludeSecretPrefixes, secret.Name) {
			continue
		}
		dataMap := make(map[string]string)
		for k := range secret.Data {
			// redact the secrets from being displayed
			dataMap[k] = "<redacted>"
		}
		secretData[secret.Name] = dataMap
	}
	if len(secretData) == 0 {
		return nil, &Error{
			Code:    http.StatusNotFound,
			Message: "No Secrets found eligible for display",
		}
	}
	return secretData, nil
}

func GetConfigMapYaml(clientset *kubernetes.Clientset, namespace string) (map[string]map[string]string, *Error) {
	configMaps, err := getconfigmapforNS(clientset, namespace)
	if err != nil {
		return nil, err
	}
	configMapData := make(map[string]map[string]string)
	for _, configMap := range configMaps.Items {
		proceed := true
		errObj := validateConfigMapEliblity(clientset, &configMap)
		//log.Printf("Error:%v\n", errObj)
		if errObj != nil {
			continue
		}
		log.Printf("ConfigMap:%s\n", configMap.Name)
		for _, name := range excludeConfigMapPrefixes {
			if strings.Contains(configMap.Name, name) {
				log.Printf("Skipping ConfigMap:%s\n", configMap.Name)
				proceed = false
				continue
			}
		}
		if !proceed {
			continue
		}
		dataMap := make(map[string]string)
		for k, v := range configMap.Data {
			dataMap[k] = string(v)
		}
		configMapData[configMap.Name] = dataMap
	}
	if len(configMapData) == 0 {
		return nil, &Error{
			Code:    http.StatusNotFound,
			Message: "No ConfigMaps found eligible for display",
		}
	}
	return configMapData, nil
}

func PatchDeploymentImage(clientset *kubernetes.Clientset, namespace string, deploymentStr string, containerName string, image string) *Error {
	deployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), deploymentStr, metav1.GetOptions{})
	if err != nil {
		return &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Only allow patching deployments that are cloned by this system using the annotations set
	errObj := validateDeploymentEliblity(clientset, deployment)
	if errObj != nil {
		return errObj
	}
	//log.Printf("Deployment: %s, Container:%s, Image:%s", deploymentStr, containerName, image)
	// Find the target container and return if not found
	containerIndex := findContainerIndex(deployment, containerName)
	if containerIndex == -1 {
		return &Error{
			Code:    http.StatusInternalServerError,
			Message: fmt.Sprintf("container %s not found in deployment %s", containerName, deployment.Name),
		}
	}

	// Check if image has already been updated to avoid unnecessary patching
	if deployment.Spec.Template.Spec.Containers[containerIndex].Image == image {
		log.Printf("Container %s in deployment %s already has image %s, skipping patch", containerName, deployment.Name, image)
		return nil
	}

	// Construct the patch to update the image
	patch := []map[string]interface{}{
		{
			"op":    "replace",
			"path":  fmt.Sprintf("/spec/template/spec/containers/%d/image", containerIndex),
			"value": image,
		},
	}

	// Marshal the patch to JSON
	patchBytes, err := json.Marshal(patch)
	if err != nil {
		return &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Try patching the deployment with retry on conflict
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		_, err := clientset.AppsV1().Deployments(namespace).Patch(context.TODO(), deployment.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{})
		return err
	})
	if retryErr != nil {
		return &Error{
			Code:    http.StatusInternalServerError,
			Message: retryErr.Error(),
		}
	}
	return nil
}

func PatchSecret(clientset *kubernetes.Clientset, namespace string, secretName string, data map[string]interface{}) *Error {

	secret, err := clientset.CoreV1().Secrets(namespace).Get(context.TODO(), secretName, metav1.GetOptions{})
	if err != nil {
		return &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	errObj := validateSecretEliblity(clientset, secret)
	if errObj != nil {
		return errObj
	}

	encodedData := make(map[string]string)
	for k, v := range data {
		strValue := fmt.Sprintf("%v", v)
		// Check if the value is a float (json numbers are treated as float64 in Go), and convert to string without scientific notation
		if floatValue, ok := v.(float64); ok {
			strValue = strconv.FormatFloat(floatValue, 'f', -1, 64)
		}
		encodedData[k] = base64.StdEncoding.EncodeToString([]byte(strValue))
	}

	// Construct the patch to update the image
	patch := []map[string]interface{}{
		{
			"op":    "replace",
			"path":  "/data",
			"value": encodedData,
		},
	}
	// Don't enable this except for debugging as this will leak all the secrets in the logs otherwise
	//log.Printf("Patch: %v", patch)

	// Marshal the patch to JSON
	patchBytes, err := json.Marshal(patch)
	if err != nil {
		return &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Try patching the deployment with retry on conflict
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		_, err := clientset.CoreV1().Secrets(namespace).Patch(context.TODO(), secret.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{})
		return err
	})
	if retryErr != nil {
		return &Error{
			Code:    http.StatusInternalServerError,
			Message: retryErr.Error(),
		}
	}
	return nil
}

func PatchConfigMap(clientset *kubernetes.Clientset, namespace string, configMapName string, data map[string]string) *Error {

	configMap, err := clientset.CoreV1().ConfigMaps(namespace).Get(context.TODO(), configMapName, metav1.GetOptions{})
	if err != nil {
		return &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	errObj := validateConfigMapEliblity(clientset, configMap)
	if errObj != nil {
		return errObj
	}

	// Construct the patch to update the image
	patch := []map[string]interface{}{
		{
			"op":    "replace",
			"path":  "/data",
			"value": data,
		},
	}
	log.Printf("Patch: %v", patch)

	// Marshal the patch to JSON
	patchBytes, err := json.Marshal(patch)
	if err != nil {
		return &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}

	// Try patching the deployment with retry on conflict
	retryErr := retry.RetryOnConflict(retry.DefaultRetry, func() error {
		_, err := clientset.CoreV1().ConfigMaps(namespace).Patch(context.TODO(), configMap.Name, types.JSONPatchType, patchBytes, metav1.PatchOptions{})
		return err
	})
	if retryErr != nil {
		log.Printf("Retry Error:%v\n", retryErr)
		return &Error{
			Code:    http.StatusInternalServerError,
			Message: retryErr.Error(),
		}
	}
	return nil
}
