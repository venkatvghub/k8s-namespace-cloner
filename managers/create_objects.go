package managers

import (
	"context"
	"fmt"
	"strings"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

func CloneConfigMap(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace string) error {
	configMaps, err := clientset.CoreV1().ConfigMaps(sourceNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Namespace doesn't have CronJobs, return successfully
			fmt.Printf("Namespace %s does not have any ConfigMaps\n", sourceNamespace)
			return nil
		} else {
			// Error checking for CronJobs
			fmt.Println("Error checking for ConfigMaps:", err)
			return err
		}
	}

	for _, configMap := range configMaps.Items {
		_, err := clientset.CoreV1().ConfigMaps(targetNamespace).Get(context.TODO(), configMap.Name, metav1.GetOptions{})
		if err != nil && !errors.IsNotFound(err) {
			// Handle unexpected errors
			fmt.Printf("Error checking for existing configmap %s: %v\n", configMap.Name, err)
			return err
		} else if err == nil {
			// ConfigMap already exists, skip creation
			fmt.Printf("ConfigMap %s already exists in %s, skipping creation\n", configMap.Name, targetNamespace)
			continue
		}
		_, err = clientset.CoreV1().ConfigMaps(targetNamespace).Create(context.TODO(), &v1.ConfigMap{
			ObjectMeta: metav1.ObjectMeta{
				Name:      configMap.Name,
				Namespace: targetNamespace,
			},
			Data: configMap.Data,
		}, metav1.CreateOptions{})

		if err != nil {
			return err
		}
		// Check if ConfigMap exists
		_, err = clientset.CoreV1().ConfigMaps(targetNamespace).Get(context.TODO(), configMap.Name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				return fmt.Errorf("ConfigMap %s not found in namespace %s", configMap.Name, targetNamespace)
			} else {
				return fmt.Errorf("Error checking for ConfigMap %s: %v", configMap.Name, err)
			}
		}

		// ConfigMap exists, return success immediately (no status to check)
		fmt.Printf("ConfigMap %s is ready\n", configMap.Name)
	}
	return nil
}

func CloneSecret(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace string) error {
	secrets, err := clientset.CoreV1().Secrets(sourceNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Namespace doesn't have CronJobs, return successfully
			fmt.Printf("Namespace %s does not have any Secrets\n", sourceNamespace)
			return nil
		} else {
			// Error checking for CronJobs
			fmt.Println("Error checking for Secrets:", err)
			return err
		}
	}

	for _, secret := range secrets.Items {
		_, err := clientset.CoreV1().Secrets(targetNamespace).Get(context.TODO(), secret.Name, metav1.GetOptions{})
		if err != nil && !errors.IsNotFound(err) {
			// Handle unexpected errors
			fmt.Printf("Error checking for existing secret %s: %v\n", secret.Name, err)
			return err
		} else if err == nil {
			// Secret already exists, skip creation
			fmt.Printf("Secret %s already exists in %s, skipping creation\n", secret.Name, targetNamespace)
			continue
		}
		_, err = clientset.CoreV1().Secrets(targetNamespace).Create(context.TODO(), &v1.Secret{
			ObjectMeta: metav1.ObjectMeta{
				Name:      secret.Name,
				Namespace: targetNamespace,
			},
			Data: secret.Data,
		}, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		// Check if Secret exists
		_, err = clientset.CoreV1().Secrets(targetNamespace).Get(context.TODO(), secret.Name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				return fmt.Errorf("Secret %s not found in namespace %s", secret.Name, targetNamespace)
			} else {
				return fmt.Errorf("Error checking for Secret %s: %v", secret.Name, err)
			}
		}

		// Secret exists, return success immediately (no status to check)
		fmt.Printf("Secret %s is ready\n", secret.Name)
	}
	return nil
}

func CloneDeployments(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace string) error {
	deployments, err := clientset.AppsV1().Deployments(sourceNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Namespace doesn't have CronJobs, return successfully
			fmt.Printf("Namespace %s does not have any Deployments\n", sourceNamespace)
			return nil
		} else {
			// Error checking for CronJobs
			fmt.Println("Error checking for Deployments:", err)
			return err
		}
	}

	for _, deployment := range deployments.Items {
		_, err := clientset.AppsV1().Deployments(targetNamespace).Get(context.TODO(), deployment.Name, metav1.GetOptions{})
		if err != nil && !errors.IsNotFound(err) {
			// Handle unexpected errors
			fmt.Printf("Error checking for existing deployment %s: %v\n", deployment.Name, err)
			return err
		} else if err == nil {
			// Deployment already exists, skip creation
			fmt.Printf("Deployment %s already exists in %s, skipping creation\n", deployment.Name, targetNamespace)
			continue
		}
		// Set desired image in container spec
		/*desiredImage := "your-desired-image" // Replace with your specific image
		for i := range deployment.Spec.Template.Spec.Containers {
			deployment.Spec.Template.Spec.Containers[i].Image = desiredImage
		}*/

		// Create deployment in target namespace
		_, err = clientset.AppsV1().Deployments(targetNamespace).Create(context.TODO(), &appsv1.Deployment{
			ObjectMeta: metav1.ObjectMeta{
				Name:      deployment.Name,
				Namespace: targetNamespace,
			},
			Spec: deployment.Spec,
		}, metav1.CreateOptions{})

		if err != nil {
			return err
		}
		// Wait for deployment to be ready
		for {
			deployment, err := clientset.AppsV1().Deployments(targetNamespace).Get(context.TODO(), deployment.Name, metav1.GetOptions{})
			if err != nil {
				return fmt.Errorf("Error getting deployment status: %v", err)
			}

			replicas := deployment.Status.ReadyReplicas
			if replicas == *(deployment.Spec.Replicas) {
				fmt.Printf("Deployment %s is ready with %d replicas\n", deployment.Name, replicas)
				return nil
			}

			// Deployment is not ready yet, check for errors
			if deployment.Status.Conditions != nil {
				for _, condition := range deployment.Status.Conditions {
					if condition.Type == appsv1.DeploymentReplicaFailure && condition.Status == corev1.ConditionTrue {
						return fmt.Errorf("Deployment %s has failed: %s", deployment.Name, condition.Reason)
					}
				}
			}

			// Deployment is still in progress, wait and try again
			time.Sleep(5 * time.Second) // Adjust the wait interval as needed
			fmt.Printf("Waiting for deployment %s to be ready (%d/%d replicas)...\n", deployment.Name, replicas, *(deployment.Spec.Replicas))
		}
		//fmt.Printf("Deployment %s cloned to %s with image %s\n", deployment.Name, targetNamespace, desiredImage)
	}
	return nil
}

func CloneServices(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace string) error {
	services, err := clientset.CoreV1().Services(sourceNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Namespace doesn't have CronJobs, return successfully
			fmt.Printf("Namespace %s does not have any Services\n", sourceNamespace)
			return nil
		} else {
			// Error checking for CronJobs
			fmt.Println("Error checking for Services:", err)
			return err
		}
	}

	for _, service := range services.Items {
		service.Spec.ClusterIP = ""           // Reset ClusterIP so that a new one is generated
		service.Spec.ClusterIPs = []string{}  // Reset ClusterIPs so that a new one is generated
		service.Spec.ExternalIPs = []string{} // Reset ExternalIPs so that a new one is generated
		service.Spec.ExternalName = ""        // Reset ExternalName so that a new one is generated
		service.Spec.LoadBalancerIP = ""      // Reset LoadBalancerIP so that a new one is generated
		_, err = clientset.CoreV1().Services(targetNamespace).Create(context.TODO(), &v1.Service{
			ObjectMeta: metav1.ObjectMeta{
				Name:      service.Name,
				Namespace: targetNamespace,
			},
			Spec: service.Spec,
		}, metav1.CreateOptions{})

		if err != nil {
			return err
		}
		for {
			// Get the latest Service status
			service, err := clientset.CoreV1().Services(targetNamespace).Get(context.TODO(), service.Name, metav1.GetOptions{})
			if err != nil {
				return fmt.Errorf("Error getting Service status: %v", err)
			}

			// Check if Service has a ClusterIP assigned (assume ready when ClusterIP is available)
			if service.Spec.ClusterIP != "" {
				fmt.Printf("Service %s is ready with ClusterIP %s\n", service.Name, service.Spec.ClusterIP)
				return nil
			}

			// Service is still being created, wait and try again
			time.Sleep(5 * time.Second) // Adjust wait interval as needed
			fmt.Printf("Waiting for Service %s to be assigned a ClusterIP...\n", service.Name)
		}
		// Service exists, return success immediately (no status to check)
	}
	return nil
}

func CloneCronJobs(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace string) error {
	cronJobs, err := clientset.BatchV1beta1().CronJobs(sourceNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Namespace doesn't have CronJobs, return successfully
			fmt.Printf("Namespace %s does not have any CronJobs\n", sourceNamespace)
			return nil
		} else {
			// Error checking for CronJobs
			fmt.Println("Error checking for CronJobs:", err)
			return err
		}
	}
	for _, cronJob := range cronJobs.Items {
		_, err = clientset.BatchV1beta1().CronJobs(targetNamespace).Create(context.TODO(), &cronJob, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		// Check if CronJob exists
		_, err := clientset.BatchV1beta1().CronJobs(targetNamespace).Get(context.TODO(), cronJob.Name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				return fmt.Errorf("CronJob %s not found in namespace %s", cronJob.Name, targetNamespace)
			} else {
				return fmt.Errorf("Error checking for CronJob %s: %v", cronJob.Name, err)
			}
		}

		// CronJob exists, return success immediately (no status to check)
		fmt.Printf("CronJob %s is ready\n", cronJob.Name)
	}
	return nil
}

func CloneJobs(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace string) error {
	jobs, err := clientset.BatchV1().Jobs(sourceNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Namespace doesn't have CronJobs, return successfully
			fmt.Printf("Namespace %s does not have any Jobs\n", sourceNamespace)
			return nil
		} else {
			// Error checking for CronJobs
			fmt.Println("Error checking for Jobs:", err)
			return err
		}
	}
	for _, job := range jobs.Items {
		_, err = clientset.BatchV1().Jobs(targetNamespace).Create(context.TODO(), &job, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		// Check if Job exists
		_, err := clientset.BatchV1().Jobs(targetNamespace).Get(context.TODO(), job.Name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				return fmt.Errorf("Job %s not found in namespace %s", job.Name, targetNamespace)
			} else {
				return fmt.Errorf("Error checking for Job %s: %v", job.Name, err)
			}
		}

		// Job exists, return success immediately (no status to check)
		fmt.Printf("Job %s is ready\n", job.Name)
	}
	return nil
}

// Helper functions for error handling
func hasStatefulSetUpdateFailure(statefulSet *appsv1.StatefulSet) bool {
	// Implement logic to check for specific failure conditions in StatefulSet status
	// Example:
	if statefulSet.Status.UpdateRevision != statefulSet.Status.CurrentRevision {
		return true
	}
	return false // Adjust based on your error detection criteria
}

func getStatefulSetFailureReason(statefulSet *appsv1.StatefulSet) string {
	// Extract the failure reason from StatefulSet status
	// Example:
	for _, condition := range statefulSet.Status.Conditions {
		if condition.Status != corev1.ConditionTrue {
			return condition.Reason
		}
	}
	return "Unknown failure reason" // Adjust based on your error reporting
}

// TODO: Need to check this
func CloneSTS(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace string) error {
	statefulSets, err := clientset.AppsV1().StatefulSets(sourceNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Namespace doesn't have CronJobs, return successfully
			fmt.Printf("Namespace %s does not have any Statefulsets\n", sourceNamespace)
			return nil
		} else {
			// Error checking for CronJobs
			fmt.Println("Error checking for Statefulsets:", err)
			return err
		}
	}
	for _, statefulSet := range statefulSets.Items {
		_, err = clientset.AppsV1().StatefulSets(targetNamespace).Create(context.TODO(), &statefulSet, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		// Check if StatefulSet exists
		_, err := clientset.AppsV1().StatefulSets(targetNamespace).Get(context.TODO(), statefulSet.Name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				return fmt.Errorf("StatefulSet %s not found in namespace %s", statefulSet.Name, targetNamespace)
			} else {
				return fmt.Errorf("Error checking for StatefulSet %s: %v", statefulSet.Name, err)
			}
		}

		// Check if StatefulSet exists
		statefulSet, err := clientset.AppsV1().StatefulSets(targetNamespace).Get(context.TODO(), statefulSet.Name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				return fmt.Errorf("StatefulSet %s not found in namespace %s", statefulSet.Name, targetNamespace)
			} else {
				return fmt.Errorf("Error checking for StatefulSet %s: %v", statefulSet.Name, err)
			}
		}

		// Wait for StatefulSet to be ready
		for {
			// Get the latest StatefulSet status
			statefulSet, err = clientset.AppsV1().StatefulSets(targetNamespace).Get(context.TODO(), statefulSet.Name, metav1.GetOptions{})
			if err != nil {
				return fmt.Errorf("Error getting StatefulSet status: %v", err)
			}

			// Check if all replicas are ready
			if statefulSet.Status.ReadyReplicas == *(statefulSet.Spec.Replicas) {
				fmt.Printf("StatefulSet %s is ready with %d replicas\n", statefulSet.Name, statefulSet.Status.ReadyReplicas)
				return nil
			}

			// Check for errors
			if hasStatefulSetUpdateFailure(statefulSet) {
				return fmt.Errorf("StatefulSet %s has failed: %s", statefulSet.Name, getStatefulSetFailureReason(statefulSet))
			}

			// StatefulSet is still rolling out, wait and try again
			time.Sleep(5 * time.Second) // Adjust the wait interval as needed
			fmt.Printf("Waiting for StatefulSet %s to be ready (%d/%d replicas)...\n", statefulSet.Name, statefulSet.Status.ReadyReplicas, *(statefulSet.Spec.Replicas))
		}
	}
	return nil
}

func CloneIngresses(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace string) error {
	ingresses, err := clientset.ExtensionsV1beta1().Ingresses(sourceNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Namespace doesn't have CronJobs, return successfully
			fmt.Printf("Namespace %s does not have any Ingress\n", sourceNamespace)
			return nil
		} else {
			// Error checking for CronJobs
			fmt.Println("Error checking for Ingress:", err)
			return err
		}
	}
	for _, ingress := range ingresses.Items {
		_, err = clientset.ExtensionsV1beta1().Ingresses(targetNamespace).Create(context.TODO(), &ingress, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		// Check if Ingress exists
		_, err := clientset.ExtensionsV1beta1().Ingresses(targetNamespace).Get(context.TODO(), ingress.Name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				return fmt.Errorf("Ingress %s not found in namespace %s", ingress.Name, targetNamespace)
			} else {
				return fmt.Errorf("Error checking for Ingress %s: %v", ingress.Name, err)
			}
		}

		// Ingress exists, return success immediately (no status to check)
		fmt.Printf("Ingress %s is ready\n", ingress.Name)
	}
	return nil
}

func ClonePDB(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace string) error {
	podDisruptionBudgets, err := clientset.PolicyV1beta1().PodDisruptionBudgets(sourceNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Namespace doesn't have CronJobs, return successfully
			fmt.Printf("Namespace %s does not have any PodDisruptionBudgets\n", sourceNamespace)
			return nil
		} else {
			// Error checking for CronJobs
			fmt.Println("Error checking for PodDisruptionBudgets:", err)
			return err
		}
	}
	for _, podDisruptionBudget := range podDisruptionBudgets.Items {
		_, err = clientset.PolicyV1beta1().PodDisruptionBudgets(targetNamespace).Create(context.TODO(), &podDisruptionBudget, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		// Check if PDB exists
		_, err := clientset.PolicyV1beta1().PodDisruptionBudgets(targetNamespace).Get(context.TODO(), podDisruptionBudget.Name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				return fmt.Errorf("PodDisruptionBudget %s not found in namespace %s", podDisruptionBudget.Name, targetNamespace)
			} else {
				return fmt.Errorf("Error checking for PodDisruptionBudget %s: %v", podDisruptionBudget.Name, err)
			}
		}

		// PDB exists, return success immediately (no status to check)
		fmt.Printf("PodDisruptionBudget %s is ready\n", podDisruptionBudget.Name)
	}
	return nil
}

func CloneHPA(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace string) error {
	horizontalPodAutoscalers, err := clientset.AutoscalingV1().HorizontalPodAutoscalers(sourceNamespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			// Namespace doesn't have CronJobs, return successfully
			fmt.Printf("Namespace %s does not have any HPA\n", sourceNamespace)
			return nil
		} else {
			// Error checking for CronJobs
			fmt.Println("Error checking for HPA:", err)
			return err
		}
	}
	for _, horizontalPodAutoscaler := range horizontalPodAutoscalers.Items {
		_, err = clientset.AutoscalingV1().HorizontalPodAutoscalers(targetNamespace).Create(context.TODO(), &horizontalPodAutoscaler, metav1.CreateOptions{})
		if err != nil {
			return err
		}
		// Check if HPA exists
		_, err := clientset.AutoscalingV1().HorizontalPodAutoscalers(targetNamespace).Get(context.TODO(), horizontalPodAutoscaler.Name, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				return fmt.Errorf("HorizontalPodAutoscaler %s not found in namespace %s", horizontalPodAutoscaler.Name, targetNamespace)
			} else {
				return fmt.Errorf("Error checking for HorizontalPodAutoscaler %s: %v", horizontalPodAutoscaler.Name, err)
			}
		}

		// HPA exists, return success immediately (no status to check)
		fmt.Printf("HorizontalPodAutoscaler %s is ready\n", horizontalPodAutoscaler.Name)
	}
	return nil
}

func RemoveNamespace(clientset *kubernetes.Clientset, namespace string) error {
	// Check if namespace exists
	_, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
	if err != nil {
		if errors.IsNotFound(err) {
			fmt.Printf("Namespace %s does not exist, nothing to delete\n", namespace)
			return nil
		} else {
			// Handle unexpected errors
			return fmt.Errorf("Error checking for namespace %s: %v", namespace, err)
		}
	}

	// Delete the namespace
	err = clientset.CoreV1().Namespaces().Delete(context.TODO(), namespace, metav1.DeleteOptions{})
	if err != nil {
		return fmt.Errorf("Error deleting namespace %s: %v", namespace, err)
	}

	// Wait for namespace deletion to complete
	for {
		_, err := clientset.CoreV1().Namespaces().Get(context.TODO(), namespace, metav1.GetOptions{})
		if err != nil {
			if errors.IsNotFound(err) {
				fmt.Printf("Namespace %s deleted successfully\n", namespace)
				return nil
			} else {
				return fmt.Errorf("Error checking for namespace deletion: %v", err)
			}
		}

		// Namespace still exists, wait and try again
		fmt.Printf("Waiting for namespace %s to be deleted...\n", namespace)
		time.Sleep(5 * time.Second) // Adjust the wait interval as needed
	}
}

func CloneNamespace(clientset *kubernetes.Clientset, sourceNamespace, targetNamespace string) error {
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

	err = CloneConfigMap(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		// Remove the Target Namespace
		// TODO: Probably move the namespace deletion to a go routine for returning faster?
		err = RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}
	err = CloneSecret(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		// Remove the Target Namespace
		err = RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = CloneDeployments(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		// Remove the Target Namespace
		err = RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = CloneServices(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		fmt.Printf("Error cloning Services: %v\n", err)
		// Remove the Target Namespace
		err = RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = CloneCronJobs(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		fmt.Printf("Error cloning CronJobs: %v\n", err)
		// Remove the Target Namespace
		err = RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = CloneJobs(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		fmt.Printf("Error cloning Jobs: %v\n", err)
		// Remove the Target Namespace
		err = RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = CloneSTS(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		fmt.Printf("Error cloning STS: %v\n", err)
		// Remove the Target Namespace
		err = RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = CloneIngresses(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		fmt.Printf("Error cloning Ingress: %v\n", err)
		// Remove the Target Namespace
		err = RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = ClonePDB(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		fmt.Printf("Error cloning PDB: %v\n", err)
		// Remove the Target Namespace
		err = RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	err = CloneHPA(clientset, sourceNamespace, targetNamespace)
	if err != nil {
		fmt.Printf("Error cloning HPA: %v\n", err)
		// Remove the Target Namespace
		err = RemoveNamespace(clientset, targetNamespace)
		if err != nil {
			fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		}
		return err
	}

	// Enable the below for testing and cleanup
	/*err = RemoveNamespace(clientset, targetNamespace)
	if err != nil {
		fmt.Errorf("Error removing namespace %s: %v\n", targetNamespace, err)
		return err
	}
	fmt.Printf("Removed Namespace:%s\n", targetNamespace)*/
	return nil
}
