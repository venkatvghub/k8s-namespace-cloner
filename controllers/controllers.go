package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	// gin-swagger middleware
	// swagger embed files
	"github.com/venkatvghub/k8s-namespace-cloner/managers"
)

type NSClonerRequestBody struct {
	//SourceNamespace string `json:"sourceNamespace"`
	TargetNamespace string `json:"targetNamespace"`
}

type DeploymentPatchRequestBody struct {
	Image string `json:"image"`
	//Deployment string `json:"deployment"`
	Container string `json:"container"`
	Namespace string `json:"namespace"`
}

type SecretPatchRequestBody struct {
	Data      map[string]interface{} `json:"data"`
	Namespace string                 `json:"namespace"`
	//SecretName string                 `json:"name"`
}

type ConfigMapPatchRequestBody struct {
	Data      map[string]string `json:"data"`
	Namespace string            `json:"namespace"`
	//ConfigMapName string            `json:"name"`
}

var (
	nsRequestBody              NSClonerRequestBody
	deploymentPatchRequestBody DeploymentPatchRequestBody
	secretPatchRequestBody     SecretPatchRequestBody
	configMapPatchRequestBody  ConfigMapPatchRequestBody
)

// GetNS godoc
// @Summary Get all namespaces
// @Description Get all namespaces in the cluster
// @Produce json
// @Success 200 {array} string
// @Router /namespaces [get]
func GetNS(c *gin.Context) {
	clientset := c.MustGet("clientset").(*kubernetes.Clientset)
	namespaceNames, err := managers.GetNS(clientset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, namespaceNames)
}

// @Summary Get deployments for a specific namespace
// @Description Get all deployments in the specified namespace
// @Produce json
// @Param namespace path string true "Namespace name"
// @Success 200 {array} string
// @Router /namespaces/:namespace/deployments [get]
func GetDeployments(c *gin.Context) {
	clientset := c.MustGet("clientset").(*kubernetes.Clientset)
	namespace := c.Param("namespace")
	deployments, err := managers.GetDeploymentForNS(clientset, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Message})
		return
	}
	if len(deployments) > 0 {
		c.JSON(http.StatusOK, deployments)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "No active deployments found"})
	}
}

// @Summary Clone a namespace
// @Description Clone a namespace and its objects to a new namespace
// @Accept json
// @Produce json
// @Param body body NSClonerRequestBody true "Namespace clone request body"
// @Success 200 {object} string
// @Router /namespaces/:namespace/cloneNamespace [post]
func CloneNamespace(c *gin.Context) {
	clientset := c.MustGet("clientset").(*kubernetes.Clientset)
	dynamicClientSet := c.MustGet("dynamicClientSet").(*dynamic.DynamicClient)
	sourceNamespace := c.Param("namespace")
	if err := c.BindJSON(&nsRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	targetNamespace := nsRequestBody.TargetNamespace
	if sourceNamespace == targetNamespace {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Source and target namespaces cannot be the same"})
		return
	}
	//sourceNamespace := nsRequestBody.SourceNamespace
	log.Printf("Source Namespace:%s, Target Namespace:%s\n", sourceNamespace, targetNamespace)
	// Implement the cloneResources function
	// Clone namespace objects
	err := managers.CloneNamespace(clientset, dynamicClientSet, sourceNamespace, targetNamespace)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Namespace %s cloned to %s. Setting Replicas to zero. When the deployment is ", sourceNamespace, targetNamespace)})
}

// @Summary Display deployments for a specific namespace
// @Description Display all deployments in the specified namespace
// @Produce json
// @Param namespace path string true "Namespace name"
// @Success 200 {object} string
// @Router /namespaces/:namespace/deployments/display [get]
func DisplayDeployments(c *gin.Context) {
	clientset := c.MustGet("clientset").(*kubernetes.Clientset)
	namespace := c.Param("namespace")
	//log.Printf("Display Deployments...")
	yamlMap, err := managers.GetDeploymentYaml(clientset, namespace)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Message})
		return
	}
	//log.Printf("Yaml MAP:%v\n", yamlMap)
	c.JSON(http.StatusOK, yamlMap)
}

// @Summary Display secrets for a specific namespace
// @Description Display all secrets in the specified namespace
// @Produce json
// @Param namespace path string true "Namespace name"
// @Success 200 {object} string
// @Router /namespaces/:namespace/secrets/display [get]
func DisplaySecrets(c *gin.Context) {
	clientset := c.MustGet("clientset").(*kubernetes.Clientset)
	namespace := c.Param("namespace")
	//log.Printf("Display Secrets...")
	yamlMap, err := managers.GetSecretYaml(clientset, namespace)
	if err != nil {
		log.Printf("Error:%v\n", err)
		c.JSON(err.Code, gin.H{"error": err.Message})
		return
	}
	//log.Printf("Yaml MAP:%v\n", yamlMap)
	c.JSON(http.StatusOK, yamlMap)
}

// @Summary Display config maps for a specific namespace
// @Description Display all config maps in the specified namespace
// @Produce json
// @Param namespace path string true "Namespace name"
// @Success 200 {object} string
// @Router /namespaces/:namespace/configmaps/display [get]
func DisplayConfigMap(c *gin.Context) {
	clientset := c.MustGet("clientset").(*kubernetes.Clientset)
	namespace := c.Param("namespace")
	//log.Printf("Display ConfigMap...")
	yamlMap, err := managers.GetConfigMapYaml(clientset, namespace)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Message})
		return
	}
	log.Printf("Yaml MAP:%v\n", yamlMap)
	c.JSON(http.StatusOK, yamlMap)
}

// @Summary Update deployment image
// @Description Update the image of a deployment in a specific namespace
// @Accept json
// @Produce json
// @Param body body DeploymentPatchRequestBody true "Deployment Image Set Request Body"
// @Success 200 {object} string
// @Router /deployments/:deployment [post]
func UpdateDeploymentImage(c *gin.Context) {
	clientset := c.MustGet("clientset").(*kubernetes.Clientset)
	deployment := c.Param("deployment")
	if err := c.BindJSON(&deploymentPatchRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	image := deploymentPatchRequestBody.Image
	//deployment := deploymentPatchRequestBody.Deployment
	container := deploymentPatchRequestBody.Container
	namespace := deploymentPatchRequestBody.Namespace

	//log.Printf("Patch Deployment...")
	err := managers.PatchDeploymentImage(clientset, namespace, deployment, container, image)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, "Patched Deployment: "+deployment+" with image: "+image)
}

// @Summary Update a secret
// @Description Update a secret in a specific namespace
// @Accept json
// @Produce json
// @Param body body SecretPatchRequestBody true "Secret Update Request Body"
// @Param secretPatchRequestBody body SecretPatchRequestBody true "Secret patch request body"
// @Success 200 {object} string
// @Router /secrets/:secret [post]
func UpdateSecret(c *gin.Context) {
	clientset := c.MustGet("clientset").(*kubernetes.Clientset)
	secretName := c.Param("secret")
	if err := c.BindJSON(&secretPatchRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	secretData := secretPatchRequestBody.Data
	namespace := secretPatchRequestBody.Namespace
	//secretName := secretPatchRequestBody.SecretName
	//log.Printf("Patch Secret...")
	err := managers.PatchSecret(clientset, namespace, secretName, secretData)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, "Patched Secret: "+secretName)
}

// @Summary Update a config map
// @Description Update a config map in a specific namespace
// @Accept json
// @Produce json
// @Param body body ConfigMapPatchRequestBody true "ConfigMap Update Request Body"
// @Success 200 {object} string
// @Router /configmaps/:configmap [post]
func UpdateConfigMap(c *gin.Context) {
	clientset := c.MustGet("clientset").(*kubernetes.Clientset)
	configMapName := c.Param("configmap")
	if err := c.BindJSON(&configMapPatchRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	configMapData := configMapPatchRequestBody.Data
	namespace := configMapPatchRequestBody.Namespace
	//configMapName := configMapPatchRequestBody.ConfigMapName
	//log.Printf("Patch ConfigMap...")
	err := managers.PatchConfigMap(clientset, namespace, configMapName, configMapData)
	if err != nil {
		c.JSON(err.Code, gin.H{"error": err.Message})
		return
	}
	c.JSON(http.StatusOK, "Patched ConfigMap: "+configMapName)

}
