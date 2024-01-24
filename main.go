package main

import (
	"flag"
	"fmt"

	"github.com/venkatvghub/k8s-ns/router"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {

	// Parse command-line arguments
	inCluster := flag.Bool("in-cluster", false, "Run inside the cluster")
	//kubeconfigPath := flag.String("kubeconfig", "", "Path to kubeconfig file")
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

	if err != nil {
		fmt.Printf("Error building Kubernetes configuration: %v\n", err)
		panic(fmt.Sprintf("Error creating Kubernetes client: %v", err))
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Printf("Error creating Kubernetes clientset: %v\n", err)
		panic(fmt.Sprintf("Error creating Kubernetes client: %v", err))
	}
	r := router.InitializeRoutes(clientset)
	fmt.Printf("Startng server at port 8080\n")
	r.Run(":8080")

}

// Utility function to get the kubeconfig path
func getKubeConfigPath() string {
	home := homedir.HomeDir()
	return home + "/.kube/config"
}
