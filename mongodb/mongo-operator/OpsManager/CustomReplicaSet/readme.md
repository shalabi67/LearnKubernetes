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