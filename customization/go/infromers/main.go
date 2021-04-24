package main

import (
	"flag"
	"fmt"
	"k8s.io/client-go/kubernetes"
	//	"context"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	api "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/tools/cache"
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
		///////////////////
		// NEVER MUTATE OBJECTS FROM INFORMERS
		/////////////////////
		AddFunc: func(new interface{}) {
			mObj := new.(metav1.Object)
			fmt.Printf("Add pod event. Pod Name= %s \n", mObj.GetName())
		},
		UpdateFunc: func(old, new interface{}) {
			mObj := new.(metav1.Object)
			pod, ok := mObj.(*api.Pod)
			if !ok {
				return
			}
			fmt.Printf("Update pod event. Pod Name= %s With status= %s\n", mObj.GetName(), pod.Status.Phase)
		},
		DeleteFunc: func(obj interface{}) {
			mObj := obj.(metav1.Object)
			fmt.Printf("Delete pod event. Pod Name= %s \n", mObj.GetName())
		},
	})
	informerFactory.Start(wait.NeverStop)
	informerFactory.WaitForCacheSync(wait.NeverStop)

	ln := ""
	fmt.Scan(&ln)
}
