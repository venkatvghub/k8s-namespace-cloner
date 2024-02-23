package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/venkatvghub/k8s-namespace-cloner/docs"
	"github.com/venkatvghub/k8s-namespace-cloner/router"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

// @title Kubernetes Namespace Cloner API
// @version 3.0.0
// @description Kubernetes Namespace Cloner API URI:<br>&nbsp;&nbsp;https://{nw-server-hostname}:8080/api/v1<br><br>
// @contact name "VV"
// @license name "No Copyrights" url "https://www.example.com"
// @host localhost:8080
// @BasePath /api/v1
// @schemes http https
func main() {

	// Parse command-line arguments
	inCluster := flag.Bool("in-cluster", false, "Run inside the cluster")
	//kubeconfigPath := flag.String("kubeconfig", "", "Path to kubeconfig file")
	// Define and parse the command line flag
	production := flag.Bool("production", false, "Start server in production mode")
	flag.Parse()

	// Initialize Kubernetes client based on the command line argument
	var config *rest.Config
	var err error

	if *inCluster {
		// Use in-cluster configuration if specified
		config, err = rest.InClusterConfig()
	} else {
		// Use external kubeconfig file if specified
		config, err = clientcmd.BuildConfigFromFlags("", getKubeConfigPath())
	}

	// Set Gin to production mode if the command line flag is specified
	if *production {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	if err != nil {
		log.Printf("Error building Kubernetes configuration: %v\n", err)
		panic(fmt.Sprintf("Error creating Kubernetes client: %v", err))
	}
	clientset, err := kubernetes.NewForConfig(config)
	//clientset, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Printf("Error creating Kubernetes clientset: %v\n", err)
		panic(fmt.Sprintf("Error creating Kubernetes client: %v", err))
	}
	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		log.Printf("Error creating Kubernetes dynamic client: %v\n", err)
		panic(fmt.Sprintf("Error creating Kubernetes client: %v", err))
	}

	r := router.InitializeRoutes(clientset, dynamicClient)
	log.Printf("Startng server at port 8080\n")
	r.Run(":8080")

}

// Utility function to get the kubeconfig path
func getKubeConfigPath() string {
	home := homedir.HomeDir()
	return home + "/.kube/config"
}
