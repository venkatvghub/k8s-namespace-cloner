package router

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/venkatvghub/k8s-namespace-cloner/controllers"
	"github.com/venkatvghub/k8s-namespace-cloner/middlewares"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
)

func InitializeRoutes(clientset *kubernetes.Clientset, dynamicClientSet *dynamic.DynamicClient) *gin.Engine {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.Use(middlewares.K8sClientSetMiddleware(clientset, dynamicClientSet))
	{
		v1.GET("/namespaces", controllers.GetNS)
		v1.GET("/namespaces/:namespace/deployments", controllers.GetDeployments)
		v1.GET("/namespaces/:namespace/deployments/display", controllers.DisplayDeployments)
		v1.GET("/namespaces/:namespace/secrets/display", controllers.DisplaySecrets)
		v1.GET("/namespaces/:namespace/configmaps/display", controllers.DisplayConfigMap)

		v1.POST("/namespaces/:namespace/cloneNamespace", controllers.CloneNamespace)
		v1.POST("/deployments/:deployment", controllers.UpdateDeploymentImage)
		v1.POST("/namespaces/:namespace/deployments/:deployment/scaleup", controllers.ScaleupDeployment)
		v1.POST("/namespaces/:namespace/deployments/:deployment/scaledown", controllers.ScaledownDeployment)
		v1.POST("/namespaces/:namespace/cronjobs/:cronjob/scaledown", controllers.ScaleDownCronjob)
		v1.POST("/namespaces/:namespace/cronjobs/:cronjob/scaleup", controllers.ScaleupCronJob)
		v1.POST("/secrets/:secret", controllers.UpdateSecret)
		v1.POST("/configmaps/:configmap", controllers.UpdateConfigMap)

	}
	// use ginSwagger middleware to serve the API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return r
}
