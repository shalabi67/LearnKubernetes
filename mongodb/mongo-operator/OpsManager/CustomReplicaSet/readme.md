# install replica set using ops manager
setup steps
- create an access key and set it in the agentConfig.yaml
- run steps in setup agent
- 

## setup
```
cd mongodb/mongo-operator/OpsManager/CustomReplicaSet/
mkdir Data
openssl rand -base64 741 > keyfile

```


## setup agent
```
kubectl apply -f replica1/agentConfig.yaml
```

## replica1
cleanup
```
sudo rm -dr Data/Replica1/*.*
sudo rm -dr Data/Replica1/journal/
sudo rm -dr Data/Replica1/WiredTiger
sudo rm Data/Logs1/mongod.log
```

setup keyfile
```
mkdir Data/Logs1
sudo chmod 777 Data
sudo chmod 777 Data/Logs1
sudo cp keyfile Data/Logs1/keyfile 
sudo chown systemd-coredump Data/Logs1/keyfile
sudo chmod 400 Data/Logs1/keyfile
```

```
kubectl label nodes ops-worker2 replica=replica1
kubectl apply -f replica1/configMap.yaml
kubectl apply -f replica1/mongod-pv.yaml
kubectl apply -f replica1/mongod-pvc.yaml
kubectl apply -f replica1/mongod-logs-pv.yaml
kubectl apply -f replica1/mongod-logs-pvc.yaml
kubectl apply -f replica1/service.yaml
kubectl apply -f replica1/service-node.yaml 
kubectl apply -f replica1/mongod.yaml
```

## replica2
cleanup
```
sudo rm -dr Data/Replica2/*.*
sudo rm -dr Data/Replica2/journal/
sudo rm -dr Data/Replica2/WiredTiger
sudo rm Data/Logs2/mongod.log
```

setup keyfile
```
mkdir Data/Logs2
sudo chmod 777 Data
sudo chmod 777 Data/Logs2
sudo cp keyfile Data/Logs2/keyfile 
sudo chown systemd-coredump Data/Logs2/keyfile
sudo chmod 400 Data/Logs2/keyfile
```

```
kubectl label nodes ops-worker3 replica=replica2
kubectl apply -f replica2/configMap.yaml
kubectl apply -f replica2/mongod-pv.yaml
kubectl apply -f replica2/mongod-pvc.yaml
kubectl apply -f replica2/mongod-logs-pv.yaml
kubectl apply -f replica2/mongod-logs-pvc.yaml
kubectl apply -f replica2/service.yaml
kubectl apply -f replica2/service-node.yaml 
kubectl apply -f replica2/mongod.yaml
```

## replica2
cleanup
```
sudo rm -dr Data/Replica3/*.*
sudo rm -dr Data/Replica3/journal/
sudo rm -dr Data/Replica3/WiredTiger
sudo rm Data/Logs3/mongod.log
```

setup keyfile
```
mkdir Data/Logs3
sudo chmod 777 Data
sudo chmod 777 Data/Logs3
sudo cp keyfile Data/Logs3/keyfile 
sudo chown systemd-coredump Data/Logs3/keyfile
sudo chmod 400 Data/Logs3/keyfile
```

```
kubectl label nodes ops-worker4 replica=replica3
kubectl apply -f replica3/configMap.yaml
kubectl apply -f replica3/mongod-pv.yaml
kubectl apply -f replica3/mongod-pvc.yaml
kubectl apply -f replica3/mongod-logs-pv.yaml
kubectl apply -f replica3/mongod-logs-pvc.yaml
kubectl apply -f replica3/service.yaml
kubectl apply -f replica3/service-node.yaml 
kubectl apply -f replica3/mongod.yaml
```
