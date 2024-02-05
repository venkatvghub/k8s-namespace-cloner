package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/venkatvghub/k8s-namespace-cloner/controllers"
	"github.com/venkatvghub/k8s-namespace-cloner/middlewares"
	"k8s.io/client-go/kubernetes"
)

func InitializeRoutes(clientset *kubernetes.Clientset) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.Use(middlewares.K8sClientSetMiddleware(clientset))
	{
		v1.GET("/namespaces", controllers.GetNS)
		v1.GET("/namespaces/:namespace/deployments", controllers.GetDeployments)
		v1.GET("/namespaces/:namespace/deployments/display", controllers.DisplayDeployments)
		v1.GET("/namespaces/:namespace/secrets/display", controllers.DisplaySecrets)
		v1.GET("/namespaces/:namespace/configmaps/display", controllers.DisplayConfigMap)

		v1.POST("/namespaces/:namespace/cloneNamespace", controllers.CloneNamespace)
		v1.POST("/deployments/:deployment/updateDeploymentImage", controllers.UpdateDeploymentImage)
		v1.POST("/secrets/:secret/updateSecret", controllers.UpdateSecret)
		v1.POST("/configmaps/:configmap/updateConfigMap", controllers.UpdateConfigMap)

	}
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
