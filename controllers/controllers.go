package controllers

import (
	"log"
        "fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// gin-swagger middleware
	// swagger embed files
	"github.com/venkatvghub/k8s-ns/managers"
	"k8s.io/client-go/kubernetes"
)

type NSClonerRequestBody struct {
	TargetNamespace string `json:"targetNamespace"`
}

var nsRequestBody NSClonerRequestBody

// @Summary List namespaces
// @Description get namespaces
// @ID get-namespaces
// @Produce json
// @Success 200 {array} string
// @Router /namespaces [get]
func GetNS(c *gin.Context) {
	clientset := c.MustGet("clientset").(*kubernetes.Clientset)
	namespaceNames, err := managers.GetNS(clientset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, namespaceNames)
}

// @Summary List deployments in a namespace
// @Description get deployments by namespace
// @ID get-deployments
// @Produce json
// @Param namespace path string true "Namespace"
// @Success 200 {array} string
// @Router /namespaces/{namespace}/deployments [get]
func GetDeployments(c *gin.Context) {
	clientset := c.MustGet("clientset").(*kubernetes.Clientset)
	namespace := c.Param("namespace")
	deployments, err := managers.GetDeploymentForNS(clientset, namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if len(deployments) > 0 {
		c.JSON(http.StatusOK, deployments)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "No active deployments found"})
	}
}

// @Summary Clone a namespace
// @Description clone namespace by source namespace
// @ID clone-namespace
// @Accept json
// @Produce json
// @Param sourceNamespace path string true "Source Namespace"
// @Success 200
// @Router /namespaces/{sourceNamespace}/clone [post]
func CloneNamespace(c *gin.Context) {
	clientset := c.MustGet("clientset").(*kubernetes.Clientset)
	sourceNamespace := c.Param("sourceNamespace")
	if err := c.BindJSON(&nsRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	targetNamespace := nsRequestBody.TargetNamespace
	log.Printf("Source Namespace:%s, Target Namespace:%s\n", sourceNamespace, targetNamespace)
	// Implement the cloneResources function
	// Clone namespace objects
	err := managers.CloneNamespace(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Namespace %s cloned to %s", sourceNamespace, targetNamespace)})
}
