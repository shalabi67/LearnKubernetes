package main

import (
	"fmt"
	"flag"
	"context"
    metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
    "k8s.io/client-go/tools/clientcmd"
    "k8s.io/client-go/kubernetes"

    "k8s.io/apimachinery/pkg/util/wait"
    "k8s.io/client-go/tools/cache"
    "k8s.io/client-go/informers"
    "time"
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

    config.AcceptContentTypes = "application/vnd.kubernetes.protobuf, application/json"
    //config.ContentType = "application/vnd.kubernetes.protobuf"

    clientset, err := kubernetes.NewForConfig(config)

    informerFactory := informers.NewSharedInformerFactory(clientset, time.Second*30)
    podInformer := informerFactory.Core().V1().Pods()
    podInformer.Informer().AddEventHandler(cache.ResourceEventHandlerFuncs{
        AddFunc: func(new interface{}) {fmt.Println("Add pod event")},
        UpdateFunc: func(old, new interface{}) {fmt.Println("Update pod event")},
        DeleteFunc: func(obj interface{}) {fmt.Println("Delete pod event")},
    })
    informerFactory.Start(wait.NeverStop)
    informerFactory.WaitForCacheSync(wait.NeverStop)

    ln := ""
    fmt.Scan(&ln)
}