package controller

import (
	"context"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/dynamic"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type InformerEvents interface {
	Add(object client.Object) error
	Update(oldObject, newObject client.Object) error
	Delete(object client.Object) error
	Convert(input interface{}) (client.Object, error)
}

func Convert(obj interface{}, object client.Object) error {
	u := obj.(*unstructured.Unstructured)
	if err := runtime.DefaultUnstructuredConverter.FromUnstructured(u.Object, object); err != nil {
		return err
	}

	return nil
}

func NewInformerEvents(informerType string, ctx context.Context, dynamicClient *dynamic.DynamicClient) InformerEvents {
	switch informerType {
	case "database":
		return &DatabaseEvents{}
	}
	return nil
}
