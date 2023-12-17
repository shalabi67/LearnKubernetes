package controller

import (
	"context"

	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	"sigs.k8s.io/controller-runtime/pkg/client"

	learnv1 "shalabi.com/informers/api/v1"
)

type DatabaseInformer struct {
	dynamicClient *dynamic.DynamicClient
	clientSet     *kubernetes.Clientset
	client        client.Client
}

func NewDatabaseInformer(dc *dynamic.DynamicClient, cs *kubernetes.Clientset, c client.Client) *DatabaseInformer {
	return &DatabaseInformer{
		dynamicClient: dc,
		clientSet:     cs,
		client:        c,
	}
}

func (d *DatabaseInformer) Build(ctx context.Context) {
	informer := NewInformerBuilder(d.dynamicClient)
	resource := learnv1.GroupVersion.WithResource("databases")
	events := NewDatabaseEvents(ctx, d.dynamicClient, d.client)
	informer.Build(ctx, resource, events)
}
