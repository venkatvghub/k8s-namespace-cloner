package managers

import (
	"context"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

type Deployment struct {
	Name      string
	Namespace string
}

func GetNS(clientset *kubernetes.Clientset) ([]string, error) {
	namespaces, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}
	namespaceNames := []string{}
	for _, namespace := range namespaces.Items {
		namespaceNames = append(namespaceNames, namespace.Name)
	}
	return namespaceNames, nil
}

func GetDeploymentForNS(clientset *kubernetes.Clientset, namespace string) ([]Deployment, error) {
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return nil, err
	}

	deploymentObjects := []Deployment{}
	for _, deployment := range deployments.Items {
		d := Deployment{}
		replicas := *deployment.Spec.Replicas
		if replicas > 0 {
			d.Name = deployment.Name
			d.Namespace = deployment.Namespace
			deploymentObjects = append(deploymentObjects, d)
		}

	}
	return deploymentObjects, nil
}
