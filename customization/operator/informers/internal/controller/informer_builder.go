package controller

import (
	"context"
	"reflect"
	"time"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/dynamic/dynamicinformer"
	"k8s.io/client-go/tools/cache"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type InformerBuilder struct {
	dynamicClient *dynamic.DynamicClient
}

func NewInformerBuilder(dynamicClient *dynamic.DynamicClient) InformerBuilder {
	return InformerBuilder{dynamicClient: dynamicClient}
}

func (i *InformerBuilder) Build(ctx context.Context, resource schema.GroupVersionResource, events InformerEvents) {
	log := log.FromContext(ctx)

	//resource := learnv1.GroupVersion.WithResource("databases")

	factory := dynamicinformer.NewFilteredDynamicSharedInformerFactory(i.dynamicClient, time.Second, "", nil)
	informer := factory.ForResource(resource).Informer()
	//cache.WaitForCacheSync(ctx.Done(), informer.HasSynced)
	//https://blog.dsb.dev/posts/creating-dynamic-informers/
	// https://fedepaol.github.io/blog/2021/01/07/writing-a-kubernetes-controller-part-2/
	informer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			o, err := events.Convert(obj)
			if err != nil {:
				log.Error(err, "AddFunc")
			}
			if err := events.Add(o); err != nil {
				log.Error(err, "events add")
			}
		},
		UpdateFunc: func(oldObj, newObj interface{}) {
			oldObject, err := events.Convert(oldObj)
			if err != nil {
				log.Error(err, "AddFunc")
			}
			newObject, err := events.Convert(newObj)
			if err != nil {
				log.Error(err, "AddFunc")
				return
			}
			if reflect.DeepEqual(oldObject, newObject) {
				log.Info("UpdateFunc object is not updated")
				return
			}

			if err := events.Update(oldObject, newObject); err != nil {
				log.Error(err, "events update")
			}
		},
		DeleteFunc: func(obj interface{}) {
			o, err := events.Convert(obj)
			if err != nil {
				log.Error(err, "AddFunc")
			}
			if err := events.Delete(o); err != nil {
				log.Error(err, "events delete")
			}
		},
	})
	informer.SetWatchErrorHandler()
	informer.Run(make(chan struct{}))
}
