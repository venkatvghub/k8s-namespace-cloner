package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venkatvghub/k8s-ns/managers"
	"k8s.io/client-go/kubernetes"
)

type NSClonerRequestBody struct {
	TargetNamespace string `json:"targetNamespace"`
}

var nsRequestBody NSClonerRequestBody

func GetNS(c *gin.Context) {
	clientset := c.MustGet("clientset").(*kubernetes.Clientset)
	namespaceNames, err := managers.GetNS(clientset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, namespaceNames)
}

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

func CloneNamespace(c *gin.Context) {
	clientset := c.MustGet("clientset").(*kubernetes.Clientset)
	sourceNamespace := c.Param("sourceNamespace")
	if err := c.BindJSON(&nsRequestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	targetNamespace := nsRequestBody.TargetNamespace
	fmt.Printf("Source Namespace:%s, Target Namespace:%s\n", sourceNamespace, targetNamespace)
	// Implement the cloneResources function
	// Clone namespace objects
	err := managers.CloneNamespace(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Namespace %s cloned to %s", sourceNamespace, targetNamespace)})
}
