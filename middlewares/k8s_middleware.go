package middlewares

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

// Middleware to inject the custom variable into the context
func K8sClientSetMiddleware(clientset *kubernetes.Clientset, dynamicClientSet *dynamic.DynamicClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("clientset", clientset)
		c.Set("dynamicClientSet", dynamicClientSet)
		c.Next()
	}
}
