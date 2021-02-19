# Pods

## create FirstPod
- kubectl create -f FirstPod.yml
- kubectl describe pods 
- kubectl get pods
- kubectl delete pods hello-pod


## busybox
- kubectl create -f busyb
- kubectl get pods
- kubectl describe pods busybox2
- kubectl get pods -o yaml
- kubectl edit pods busybox2
- kubectl delete pods busybox2


## multi container
- kubectl create -f multi-containers.yml
- kubectl describe pods busybox
- kubectl delete pods busybox


## multi containers to support logging
- kubectl create -f multi-containers\sidecar.yml
- kubectl create -f multi-containers\sidecar-service.yml
- http://172.17.200.167:30002/date.txt

### connect to the sidecar container
- kubectl exec -it sidecar-pod -c sidecar /bin/bash

#### execute the following in the sidecar bash
- 


### clean up
- kubectl delete -f multi-containers\sidecar.yml
- kubectl delete -f multi-containers\sidecar-service.yml

## create pod on a namespace
- kubectl create ns testing
- kubectl create -f namespace.yaml
- kubectl get pods -n testing
- kubectl describe pods busybox2 -n testing

### you can use the explain command to find how to identify the name space
- kubectl explain pods
- kubectl explain pods.metadata





