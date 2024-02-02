package managers

import (
	"context"
	"log"

	"net/http"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

var errorCodes = map[string]int{
	"NamespaceAnnotationMissing":  http.StatusBadRequest,
	"DeploymentAnnotationMissing": http.StatusBadRequest,
	"ConfigMapAnnotationMissing":  http.StatusBadRequest,
	"SecretAnnotationMissing":     http.StatusBadRequest,
}

type Error struct {
	Code    int
	Message string
}

func validateSourceNamespace(clientset *kubernetes.Clientset, sourceNamespace string) *Error {
	namespace, err := clientset.CoreV1().Namespaces().Get(context.TODO(), sourceNamespace, metav1.GetOptions{})
	if err != nil {
		return &Error{
			Code:    http.StatusInternalServerError,
			Message: err.Error(),
		}
	}
	annotations := namespace.Annotations
	if annotations != nil {
		if _, ok := annotations[NS_CLONER_ANNOTATION]; ok {
			//log.Printf("Annotations:%v\n", annotations[NS_CLONER_ANNOTATION])
			if !(annotations[NS_CLONER_ANNOTATION] == "true" || annotations[NS_CLONER_ANNOTATION] == "True") {
				return &Error{
					Code:    errorCodes["NamespaceAnnotationMissing"],
					Message: "Source namespace is not cloneable",
				}
			}
		} else {
			return &Error{
				Code:    errorCodes["NamespaceAnnotationMissing"],
				Message: "Source namespace is not cloneable",
			}
		}
	}
	return nil
}

func validateDeploymentEliblity(clientset *kubernetes.Clientset, deployment *appsv1.Deployment) *Error {
	// Check if the deployment is already cloned
	annotations := deployment.Annotations
	if annotations != nil {
		if _, ok := annotations[TARGET_NS_ANNOTATION_ENABLED]; ok {
			//log.Printf("Annotations:%v\n", annotations[NS_CLONER_ANNOTATION])
			if !(annotations[TARGET_NS_ANNOTATION_ENABLED] == "true" || annotations[TARGET_NS_ANNOTATION_ENABLED] != "True") {
				return &Error{
					Code:    errorCodes["DeploymentAnnotationMissing"],
					Message: "Deployment is not Annotated for operations",
				}
			}
		} else {
			return &Error{
				Code:    errorCodes["DeploymentAnnotationMissing"],
				Message: "Deployment is not Annotated for operations",
			}
		}
	}
	return nil
}

func validateSecretEliblity(clientset *kubernetes.Clientset, secret *v1.Secret) *Error {
	// Check if the deployment is already cloned
	annotations := secret.ObjectMeta.Annotations
	if annotations != nil {
		if _, ok := annotations[TARGET_NS_ANNOTATION_ENABLED]; ok {
			//log.Printf("Annotations:%v\n", annotations[NS_CLONER_ANNOTATION])
			if !(annotations[TARGET_NS_ANNOTATION_ENABLED] == "true" || annotations[TARGET_NS_ANNOTATION_ENABLED] != "True") {
				return &Error{
					Code:    errorCodes["SecretAnnotationMissing"],
					Message: "Secret is not Annotated for Operations",
				}
			}
		} else {
			return &Error{
				Code:    errorCodes["SecretAnnotationMissing"],
				Message: "Secret is not Annotated for Operations",
			}
		}
	}
	return nil
}

func validateConfigMapEliblity(clientset *kubernetes.Clientset, configMap *v1.ConfigMap) *Error {
	// Check if the deployment is already cloned
	//annotations := configMap.Annotations
	annotations := configMap.ObjectMeta.Annotations
	//log.Printf("Config Map:%s\n", configMap.Name)
	if annotations != nil {
		if _, ok := annotations[TARGET_NS_ANNOTATION_ENABLED]; ok {
			//log.Printf("Annotations:%v\n", annotations[NS_CLONER_ANNOTATION])
			if !(annotations[TARGET_NS_ANNOTATION_ENABLED] == "true" || annotations[TARGET_NS_ANNOTATION_ENABLED] != "True") {
				log.Printf("Namespace annotated")
				return &Error{
					Code:    errorCodes["ConfigMapAnnotationMissing"],
					Message: "ConfigMap is not Annotated for operations",
				}
			}
		} else {
			log.Printf("Namespace Not annotated")
			return &Error{
				Code:    errorCodes["ConfigMapAnnotationMissing"],
				Message: "ConfigMap is not Annotated for operations",
			}
		}
	}
	return nil
}
