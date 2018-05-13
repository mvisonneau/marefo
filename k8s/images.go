package k8s

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

// List all the pods currently running in K8S and return their images
func FetchImagesFromRunningPods() (i map[string]interface{}) {
	// Initialize the output variable
	i = make(map[string]interface{})
	i["images"] = make([]string, 0)

	// We assume that we run withing a K8S cluster
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// Initiate the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	// List running pods in all namespaces
	pods, err := clientset.CoreV1().Pods("").List(metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}

	// Gather container images from running pods
	for _, pod := range pods.Items {
		for _, container := range pod.Spec.Containers {
			i["images"] = appendIfMissing(i["images"].([]string), container.Image)
		}
	}

	return
}

// Function in order to only append a string to a sliace if the string doesn't already exist
func appendIfMissing(slice []string, str string) []string {
	for _, ele := range slice {
		if ele == str {
			return slice
		}
	}
	return append(slice, str)
}
