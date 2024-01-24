package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/venkatvghub/k8s-ns/managers"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

type NSClonerRequestBody struct {
	TargetNamespace string `json:"targetNamespace"`
}

var nsRequestBody NSClonerRequestBody

func main() {
	r := gin.Default()
	// Initialize Kubernetes client
	config, err := clientcmd.BuildConfigFromFlags("", getKubeConfigPath())
	if err != nil {
		panic(fmt.Sprintf("Error building kubeconfig: %v", err))
	}
	//fmt.Printf("Kube Config:%v\n", config)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(fmt.Sprintf("Error creating Kubernetes client: %v", err))
	}
	fmt.Printf("Startng server at port 8080\n")

	// API endpoint to discover deployments in a namespace
	r.GET("/api/discover/:namespace/deployments", func(c *gin.Context) {
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
	})

	// API endpoint to discover namespaces
	r.GET("/api/namespaces", func(c *gin.Context) {
		namespaceNames, err := managers.GetNS(clientset)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, namespaceNames)
	})
	// API endpoint to clone resources to a target namespace
	r.POST("/api/clone/:sourceNamespace", func(c *gin.Context) {
		sourceNamespace := c.Param("sourceNamespace")
		if err := c.BindJSON(&nsRequestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		targetNamespace := nsRequestBody.TargetNamespace
		fmt.Printf("Source Namespace:%s, Target Namespace:%s\n", sourceNamespace, targetNamespace)
		// Implement the cloneResources function
		// Clone namespace objects
		err := cloneNamespace(clientset, sourceNamespace, targetNamespace)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Namespace %s cloned to %s", sourceNamespace, targetNamespace)})
	})
	// Run the server
	r.Run(":8080")
}

// Utility function to get the kubeconfig path
func getKubeConfigPath() string {
	home := homedir.HomeDir()
	return home + "/.kube/config"
}

func cloneNamespace(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace string) error {
	// Create the target namespace if it doesn't exist
	_, err := clientset.CoreV1().Namespaces().Create(context.TODO(), &v1.Namespace{
		ObjectMeta: metav1.ObjectMeta{
			Name: targetNamespace,
		},
	}, metav1.CreateOptions{})

	if err != nil && !strings.Contains(err.Error(), "AlreadyExists") {
		fmt.Printf("Error creating namespace %s: %v\n", targetNamespace, err)
		return err
	}

	err = managers.CloneConfigMap(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		// Remove the Target Namespace
		// TODO: Probably move the namespace deletion to a go routine for returning faster?
		err = managers.RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}
	err = managers.CloneSecret(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		// Remove the Target Namespace
		err = managers.RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = managers.CloneDeployments(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		// Remove the Target Namespace
		err = managers.RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = managers.CloneServices(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		fmt.Printf("Error cloning Services: %v\n", err)
		// Remove the Target Namespace
		err = managers.RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = managers.CloneCronJobs(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		fmt.Printf("Error cloning CronJobs: %v\n", err)
		// Remove the Target Namespace
		err = managers.RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = managers.CloneJobs(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		fmt.Printf("Error cloning Jobs: %v\n", err)
		// Remove the Target Namespace
		err = managers.RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = managers.CloneSTS(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		fmt.Printf("Error cloning STS: %v\n", err)
		// Remove the Target Namespace
		err = managers.RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = managers.CloneIngresses(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		fmt.Printf("Error cloning Ingress: %v\n", err)
		// Remove the Target Namespace
		err = managers.RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = managers.ClonePDB(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		fmt.Printf("Error cloning PDB: %v\n", err)
		// Remove the Target Namespace
		err = managers.RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = managers.CloneHPA(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		fmt.Printf("Error cloning HPA: %v\n", err)
		// Remove the Target Namespace
		err = managers.RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	// Enable the below for testing and cleanup
	/*err = managers.RemoveNamespace(clientset, targetNamespace)
	if err != nil {
		fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		return err
	}
	fmt.Printf("Removed Namespace:%s\n", targetNamespace)*/
	return nil
}
