package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/venkatvghub/k8s-ns/controllers"
	"github.com/venkatvghub/k8s-ns/middlewares"
	"k8s.io/client-go/kubernetes"
)

func InitializeRoutes(clientset *kubernetes.Clientset) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.Use(middlewares.K8sClientSetMiddleware(clientset))
	{
		v1.GET("/namespaces", controllers.GetNS)
		v1.GET("/namespaces/:namespace/deployments", controllers.GetDeployments)
		v1.POST("/namespaces/:sourceNamespace/clone", controllers.CloneNamespace)
	}
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", func(c *gin.Context) {
		http.StripPrefix("/swagger/", http.FileServer(http.Dir("docs"))).ServeHTTP(c.Writer, c.Request)
	})
	return r
}
