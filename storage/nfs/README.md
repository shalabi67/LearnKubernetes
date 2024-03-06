# NFS
This example shows how to use NFS in Kind environment. 
This will be used as a basic to learn PV and PVC.
[kubernetes example code](https://github.com/kubernetes/examples/tree/master/staging/volumes/nfs)


## Setup
### Setup kind
[example to delete](https://stackoverflow.com/questions/62694361/how-to-reference-a-local-volume-in-kind-kubernetes-in-docker)
```shell
cd storage/kind
export KUBECONFIG=~/.kube/kind-nfs
kind create cluster --config kind-config.yaml --name nfs
kubectl apply -f pv.yaml 
kubectl apply -f pvc.yaml 
kubectl apply -f nfs-server.yaml
kubectl apply -f nfs-service.yaml
```

The test-yaml app will test if the hostpath works on kind.
```shell
kubectl apply -f test.yaml
```

### Use nfs volume
```shell
kubectl apply -f nfs-pv.yaml
kubectl apply -f nfs-pvc.yaml
kubectl apply -f web-app.yaml 

```