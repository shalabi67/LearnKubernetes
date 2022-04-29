# use mongodb 
https://devopscube.com/deploy-mongodb-kubernetes/


## kind setup
```
kind delete cluster
kind create cluster --config ../kind/kind-config.yaml
```

## Mongo setup
```
 export KUBECONFIG=~/.kube/kind
 kubectl apply -f configMap.yaml 
 kubectl apply -f mongod-pv.yaml
 kubectl apply -f mongod-pvc.yaml 
 kubectl apply -f mongod.yaml
 kubectl apply -f mongo.yaml
```

to create a config map
```
kubectl create cm nginy-server-config --from-file config.yaml
```

## Test setup
```
export KUBECONFIG=~/.kube/kind
kubectl exec -it mongo-client -- /bin/bash
mongo
show dbs
```

