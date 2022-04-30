# replica sets

## setup kind
```
kind delete cluster
cd ReplicaSet
kind create cluster --config kind/kind-config.yaml
kubectl label nodes kind-worker replica=replica1
kubectl label nodes kind-worker2 replica=replica2
kubectl label nodes kind-worker3 replica=replica3
sudo chmod 777 Data/Logs1
sudo chmod 777 Data/Logs2
sudo chmod 777 Data/Logs3

sudo cp keyfile Data/Logs1/keyfile 
sudo chown systemd-coredump Data/Logs1/keyfile
sudo chmod 400 Data/Logs1/keyfile

sudo cp keyfile Data/Logs2/keyfile 
sudo chown systemd-coredump Data/Logs2/keyfile
sudo chmod 400 Data/Logs2/keyfile

sudo cp keyfile Data/Logs3/keyfile 
sudo chown systemd-coredump Data/Logs3/keyfile
sudo chmod 400 Data/Logs3/keyfile

sudo rm Data/Logs1/mongod.log
sudo rm Data/Logs2/mongod.log
sudo rm Data/Logs3/mongod.log
```
## setup
create keyfile used to authenticate between replica sets
```
openssl rand -base64 741 > keyfile
chmod 600 keyfile
kubectl apply -f keyFile-configMap.yaml
```

### replica1
```
kubectl apply -f replica1/configMap.yaml
kubectl apply -f replica1/mongod-pv.yaml
kubectl apply -f replica1/mongod-pvc.yaml
kubectl apply -f replica1/mongod-logs-pv.yaml
kubectl apply -f replica1/mongod-logs-pvc.yaml
kubectl apply -f replica1/service.yaml
kubectl apply -f replica1/mongod.yaml
```


### replica2
```
kubectl apply -f replica2/configMap.yaml
kubectl apply -f replica2/mongod-pv.yaml
kubectl apply -f replica2/mongod-pvc.yaml
kubectl apply -f replica2/mongod-logs-pv.yaml
kubectl apply -f replica2/mongod-logs-pvc.yaml
kubectl apply -f replica2/service.yaml
kubectl apply -f replica2/mongod.yaml
```


### replica3
```
kubectl apply -f replica3/configMap.yaml
kubectl apply -f replica3/mongod-pv.yaml
kubectl apply -f replica3/mongod-pvc.yaml
kubectl apply -f replica3/mongod-logs-pv.yaml
kubectl apply -f replica3/mongod-logs-pvc.yaml
kubectl apply -f replica3/service.yaml
kubectl apply -f replica3/mongod.yaml
```

## configure replica
```
mongo --port 27011
rs.initiate()
use admin
db.createUser({
    user: "mohammad",
    pwd: "password123",
    roles: [
      {role: "root", db: "admin"}
    ]
  })

mongo --host "repl-example/localhost:27011" --username mohammad --password password123 --authenticationDatabase admin
rs.status()
rs.add("replica2:27012")
rs.add("replica3:27013")
rs.isMaster()
```