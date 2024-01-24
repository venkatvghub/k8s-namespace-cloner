package middlewares

import (
	"github.com/gin-gonic/gin"
	"k8s.io/client-go/kubernetes"
)

// Middleware to inject the custom variable into the context
func K8sClientSetMiddleware(clientset *kubernetes.Clientset) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("clientset", clientset)
		c.Next()
	}
}
