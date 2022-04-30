# use mongodb 
https://devopscube.com/deploy-mongodb-kubernetes/


## kind setup
```
kind delete cluster --name single
kind create cluster --name single --config ../kind/kind-config.yaml
kubectl config set-context kind-single

sudo rm -dr ../data/mongo/*.*
sudo rm -dr ../data/mongo/logs
sudo rm -dr ../data/mongo/journal/
sudo rm -dr ../data/mongo/WiredTiger
```

## Mongo setup
```
 export KUBECONFIG=~/.kube/kind
 kubectl apply -f configMap.yaml 
 kubectl apply -f mongod-pv.yaml
 kubectl apply -f mongod-pvc.yaml 
 kubectl apply -f mongod.yaml
```

to create a config map
```
kubectl create cm nginy-server-config --from-file config.yaml
```

## Test setup
```
export KUBECONFIG=~/.kube/kind
kubectl exec -it mongod-1 -- /bin/bash
mongo
show dbs

use admin
db.stats()
```
### security
```
use admin
db.createUser({
    user: "mohammad",
    pwd: "password123",
    roles: [
      {role: "root", db: "admin"}
    ]
  })
mongo --username mohammad --password password123 --authenticationDatabase admin
```

### find mongo applications
```
find /usr/bin/ -name "mongo*"
```

### use mongo tools
copy data to pod
```
kubectl cp ../databases/products.json mongod-1:/data/products.json

```

importing data
```
mongoimport -d shop -c products -u mohammad -p password123 --authenticationDatabase admin /data/products.json
mongo --username mohammad --password password123 --authenticationDatabase admin
use shop
db.products.find()
```

