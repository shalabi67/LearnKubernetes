package main

import (
	"fmt"
	"flag"
	"context"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/kubernetes"
)

func main() {
	fmt.Println("Get Pods")

	kubeconfig := flag.String("kubeconfig", "/home/mohammad/.kube/config", "kubeconfig file")
    flag.Parse()

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
    if err != nil {
        fmt.Printf("The kubeconfig cannot be loaded: %v\n", err)
        return
    }
    flag.Parse()
    config, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
    clientset, err := kubernetes.NewForConfig(config)

    namespace := "default"
    podName := "go-pod"
    pod, err := clientset.CoreV1().Pods(namespace).Get(context.TODO(), podName, metav1.GetOptions{})
    if err!=nil {
        fmt.Printf("Pod not found: %v\n", err)
        return
    }

    fmt.Println(pod);
}