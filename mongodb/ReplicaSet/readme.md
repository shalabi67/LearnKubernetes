# replica sets

## setup kind
```
kind delete cluster
cd ReplicaSet
kind create cluster --config kind/kind-config.yaml
kubectl label nodes kind-worker replica=replica1
kubectl label nodes kind-worker2 replica=replica2
kubectl label nodes kind-worker3 replica=replica3
kubectl label nodes kind-worker4 replica=arbiter
kubectl label nodes kind-worker5 replica=hidden
kubectl label nodes kind-worker6 replica=slave-delay
sudo chmod 777 Data/Logs1
sudo chmod 777 Data/Logs2
sudo chmod 777 Data/Logs3
sudo chmod 777 Data/Logs4
sudo chmod 777 Data/Logs5
sudo chmod 777 Data/Logs6

sudo cp keyfile Data/Logs1/keyfile 
sudo chown systemd-coredump Data/Logs1/keyfile
sudo chmod 400 Data/Logs1/keyfile

sudo cp keyfile Data/Logs2/keyfile 
sudo chown systemd-coredump Data/Logs2/keyfile
sudo chmod 400 Data/Logs2/keyfile

sudo cp keyfile Data/Logs3/keyfile 
sudo chown systemd-coredump Data/Logs3/keyfile
sudo chmod 400 Data/Logs3/keyfile

sudo cp keyfile Data/Logs4/keyfile 
sudo chown systemd-coredump Data/Logs4/keyfile
sudo chmod 400 Data/Logs4/keyfile

sudo cp keyfile Data/Logs5/keyfile 
sudo chown systemd-coredump Data/Logs5/keyfile
sudo chmod 400 Data/Logs5/keyfile

sudo cp keyfile Data/Logs6/keyfile 
sudo chown systemd-coredump Data/Logs6/keyfile
sudo chmod 400 Data/Logs6/keyfile

sudo rm Data/Logs1/mongod.log
sudo rm Data/Logs2/mongod.log
sudo rm Data/Logs3/mongod.log
sudo rm Data/Logs4/mongod.log
sudo rm Data/Logs5/mongod.log
sudo rm Data/Logs6/mongod.log
```
## setup
create keyfile used to authenticate between replica sets
```
openssl rand -base64 741 > keyfile
chmod 600 keyfile
kubectl apply -f keyFile-configMap.yaml
```

### replica1
cleanup
```
sudo rm -dr Data/Replica1/*.*
sudo rm -dr Data/Replica1/journal/
sudo rm -dr Data/Replica1/WiredTiger
sudo rm Data/Logs1/mongod.log
```

```
kubectl apply -f replica1/configMap.yaml
kubectl apply -f replica1/mongod-pv.yaml
kubectl apply -f replica1/mongod-pvc.yaml
kubectl apply -f replica1/mongod-logs-pv.yaml
kubectl apply -f replica1/mongod-logs-pvc.yaml
kubectl apply -f replica1/service.yaml
kubectl apply -f replica1/mongod.yaml

mongo --disableImplicitSessions --port 27011 --eval "db.adminCommand('ping')" 

mongo --host "repl-example/replica1:27011" --username mohammad --password password123 --authenticationDatabase admin

```


### replica2
cleanup
```
sudo rm -dr Data/Replica2/*.*
sudo rm -dr Data/Replica2/journal/
sudo rm -dr Data/Replica2/WiredTiger
sudo rm Data/Logs2/mongod.log
```

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
cleanup
```
sudo rm -dr Data/Replica3/*.*
sudo rm -dr Data/Replica3/journal/
sudo rm -dr Data/Replica3/WiredTiger
sudo rm Data/Logs3/mongod.log
```

```
kubectl apply -f replica3/configMap.yaml
kubectl apply -f replica3/mongod-pv.yaml
kubectl apply -f replica3/mongod-pvc.yaml
kubectl apply -f replica3/mongod-logs-pv.yaml
kubectl apply -f replica3/mongod-logs-pvc.yaml
kubectl apply -f replica3/service.yaml
kubectl apply -f replica3/mongod.yaml
```

### arbiter (replica4)
cleanup
```
sudo rm -dr Data/Replica4/*.*
sudo rm -dr Data/Replica4/journal/
sudo rm -dr Data/Replica4/WiredTiger
sudo rm Data/Logs4/mongod.log
```

```
kubectl apply -f arbiter/configMap.yaml
kubectl apply -f arbiter/mongod-pv.yaml
kubectl apply -f arbiter/mongod-pvc.yaml
kubectl apply -f arbiter/mongod-logs-pv.yaml
kubectl apply -f arbiter/mongod-logs-pvc.yaml
kubectl apply -f arbiter/service.yaml
kubectl apply -f arbiter/mongod.yaml
```

### hidden (replica5)
cleanup
```
sudo rm -dr Data/Replica5/*.*
sudo rm -dr Data/Replica5/journal/
sudo rm -dr Data/Replica5/WiredTiger
sudo rm Data/Logs5/mongod.log
```

```
kubectl apply -f hidden/configMap.yaml
kubectl apply -f hidden/mongod-pv.yaml
kubectl apply -f hidden/mongod-pvc.yaml
kubectl apply -f hidden/mongod-logs-pv.yaml
kubectl apply -f hidden/mongod-logs-pvc.yaml
kubectl apply -f hidden/service.yaml
kubectl apply -f hidden/mongod.yaml
```

### hidden (replica5)
cleanup
```
sudo rm -dr Data/Replica5/*.*
sudo rm -dr Data/Replica5/journal/
sudo rm -dr Data/Replica5/WiredTiger
sudo rm Data/Logs5/mongod.log
```

```
kubectl apply -f hidden/configMap.yaml
kubectl apply -f hidden/mongod-pv.yaml
kubectl apply -f hidden/mongod-pvc.yaml
kubectl apply -f hidden/mongod-logs-pv.yaml
kubectl apply -f hidden/mongod-logs-pvc.yaml
kubectl apply -f hidden/service.yaml
kubectl apply -f hidden/mongod.yaml
```

### slave-delay (replica6)
cleanup
```
sudo rm -dr Data/Replica6/*.*
sudo rm -dr Data/Replica6/journal/
sudo rm -dr Data/Replica6/WiredTiger
sudo rm Data/Logs6/mongod.log
```

```
kubectl apply -f slave-delay/configMap.yaml
kubectl apply -f slave-delay/mongod-pv.yaml
kubectl apply -f slave-delay/mongod-pvc.yaml
kubectl apply -f slave-delay/mongod-logs-pv.yaml
kubectl apply -f slave-delay/mongod-logs-pvc.yaml
kubectl apply -f slave-delay/service.yaml
kubectl apply -f slave-delay/mongod.yaml
```


## configure replica
```
mongosh
rs.initiate()
use admin
db.createUser({
    user: "mohammad",
    pwd: "password123",
    roles: [
      {role: "root", db: "admin"}
    ]
  })

mongosh --host "repl-example/localhost:27017" --username mohammad --password password123 --authenticationDatabase admin
rs.status()
rs.add("replica2:27017")
rs.add("replica3:27017")
rs.isMaster()
rs.stepDown()
rs.status()
```

Add arbiter
```
db.adminCommand(
  {
    setDefaultRWConcern : 1,
    defaultReadConcern: { "level": "majority" },
    defaultWriteConcern: { w : "majority" },
  }
)
rs.add( {host: "arbiter:27017", arbiterOnly: true } )
rs.status()
```

Add hidden
```
rs.add( {host: "hidden:27017", hidden: true, priority: 0 } )
rs.status()
```

Add slave-delay
```
rs.add( {host: "slave-delay:27017", hidden: true, priority: 0, secondaryDelaySecs:3600, votes: 0 } )
rs.status()
```