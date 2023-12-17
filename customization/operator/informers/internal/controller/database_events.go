package controller

import (
	"context"
	"fmt"

	v1 "k8s.io/api/core/v1"
	v1api "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/dynamic"
	learnv1 "shalabi.com/informers/api/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

type DatabaseEvents struct {
	ctx           context.Context
	dynamicClient *dynamic.DynamicClient
	client        client.Client
}

func NewDatabaseEvents(ctx context.Context, dynamicClient *dynamic.DynamicClient, client client.Client) *DatabaseEvents {
	return &DatabaseEvents{
		ctx:           ctx,
		dynamicClient: dynamicClient,
		client:        client,
	}
}

func (d *DatabaseEvents) Add(object client.Object) error {
	log := log.FromContext(d.ctx)
	database := object.(*learnv1.Database)
	log.Info(fmt.Sprintf("AddFunc %v", database))

	nameSpace := &v1.Namespace{
		ObjectMeta: v1api.ObjectMeta{Name: database.Name},
	}
	if err := d.client.Create(d.ctx, nameSpace); err != nil {
		log.Error(err, "could not create server")
		return err
	}

	for i := 0; i < database.Spec.Replicas; i++ {
		if err := d.createServer(database, i+1); err != nil {
			log.Error(err, "could not create server")
			return err
		}
	}
	return nil
}
func (d *DatabaseEvents) Update(oldObject, newObject client.Object) error {
	log := log.FromContext(d.ctx)
	log.Info("UpdateFunc")
	return nil
}
func (d *DatabaseEvents) Delete(object client.Object) error {
	log := log.FromContext(d.ctx)
	log.Info("DeleteFunc")
	return nil
}
func (d *DatabaseEvents) Convert(input interface{}) (client.Object, error) {
	var db learnv1.Database
	if err := Convert(input, &db); err != nil {
		return nil, err
	}
	return &db, nil
}

func (d *DatabaseEvents) createServer(database *learnv1.Database, index int) error {
	server := &learnv1.Server{
		ObjectMeta: v1api.ObjectMeta{
			Name:      fmt.Sprintf("%s-%d", database.Name, index),
			Namespace: database.Name,
		},
		Spec: learnv1.ServerSpec{
			CPU:     database.Spec.CPU,
			Memory:  database.Spec.Memory,
			Storage: database.Spec.Storage,
		},
	}

	if err := d.client.Create(d.ctx, server); err != nil {
		return err
	}

	return nil
}
