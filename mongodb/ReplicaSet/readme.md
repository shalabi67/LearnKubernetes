# replica sets

## setup host
```
sudo cp /etc/hosts hosts
sudo chown mshalabai hosts
echo "# The following lines for mongo replication" >>hosts
echo "127.0.0.1	  replica1">>hosts
echo "127.0.0.1	  replica2">>hosts
echo "127.0.0.1	  replica3">>hosts
sudo chown root hosts
sudo mv hosts /etc/hosts 
```
## setup kind
```
export KUBECONFIG=~/.kube/kind 
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
kubectl apply -f replica1/service-node.yaml 
kubectl apply -f replica1/mongod.yaml

mongo --disableImplicitSessions --port 27017 --eval "db.adminCommand('ping')" 

mongo --host "repl-example/replica1:27017" --username mohammad --password password123 --authenticationDatabase admin
```

connect from host machine
```
mongo --host "localhost:27001" --username mohammad --password password123 --authenticationDatabase admin
mongo --host "repl-example/replica1:27001" --username mohammad --password password123 --authenticationDatabase admin
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
kubectl apply -f replica2/service-node.yaml 
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
kubectl apply -f replica3/service-node.yaml 
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
kubectl apply -f arbiter/service-node.yaml 
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
kubectl apply -f hidden/service-node.yaml 
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
kubectl apply -f slave-delay/service-node.yaml 
kubectl apply -f slave-delay/mongod.yaml
```


## configure replica
```
mongosh --port 27001
rs.initiate()
use admin
db.createUser({
    user: "mohammad",
    pwd: "password123",
    roles: [
      {role: "root", db: "admin"}
    ]
  })

mongosh --host "repl-example/localhost:27001" --username mohammad --password password123 --authenticationDatabase admin
rs.status()
rs.add("replica2:27002")
rs.add("replica3:27003")
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
rs.add( {host: "arbiter:27004", arbiterOnly: true } )
rs.status()
```

Add hidden
```
rs.add( {host: "hidden:27005", hidden: true, priority: 0 } )
rs.status()
```

Add slave-delay
```
rs.add( {host: "slave-delay:27005", hidden: true, priority: 0, secondaryDelaySecs:3600, votes: 0 } )
rs.status()
```

## read from secondary
initialize data
```
mongosh --host "repl-example/localhost:27001" --username mohammad --password password123 --authenticationDatabase admin
use newDB
db.new_collection.insertOne( { "student": "Matt Javaly", "grade": "A+" } )
```

read from secondary
```
mongosh --host "replica2:27002" --username mohammad --password password123 --authenticationDatabase admin
use newDB
rs.slaveOk() or db.getMongo().setReadPref("primaryPreferred")
db.new_collection.find()
```

write to secondary should fail
```
use newDB
db.new_collection.insertOne( { "student": "Matt Javaly", "grade": "A+" } )
```

## oplog
```
use local
show collections
```

It is cap collection
```
var stats = db.oplog.rs.stats()
stats.capped
stats.size
stats.maxSize
```

show size in GB
```
var stats = db.oplog.rs.stats(1024*1024*1024)
stats.maxSize

df -h
```

show size in MB
```
var stats = db.oplog.rs.stats(1024*1024)
stats.maxSize
```

rs.printReplicationInfo()
```
rs.printReplicationInfo()
```

### single command can end up with multiple entries in the oplog
```
use demo
db.createCollection('messages')
```

Query the oplog, filtering out the heartbeats ("periodic noop") and only returning the latest entry:
```
use local
db.oplog.rs.find( { "o.msg": { $ne: "periodic noop" } } ).sort( { $natural: -1 } ).limit(1).pretty()
```

Insert 100 different documents:
```
use demo
for ( i=0; i< 100; i++) { db.messages.insert( { 'msg': 'not yet', _id: i } ) }
db.messages.count()
```

Query the oplog to find all operations related to demo.messages:
```
use local
db.oplog.rs.find({"ns": "demo.messages"}).sort({$natural: -1})
```

Illustrate that one update statement may generate many entries in the oplog:
```
use demo
db.messages.updateMany( {}, { $set: { author: 'norberto' } } )
use local
db.oplog.rs.find( { "ns": "demo.messages" } ).sort( { $natural: -1 } )
```

## reconfigure node
This will configure the slave-delay node to be just hidden
```
cfg = rs.conf()
cfg.members[5].votes = 0
cfg.members[5].hidden = true
cfg.members[5].priority = 0
cfg.members[5].secondaryDelaySecs = 0
rs.reconfig(cfg)
rs.conf()
```

bring the node back to be slave-delay with one hour delay
```
cfg = rs.conf()
cfg.members[5].secondaryDelaySecs = 3600
rs.reconfig(cfg)
rs.conf()
```

## commands
```
db.new_collection.insertOne( { "student": "Matt Javaly", "grade": "A+" } , {writeConcern: {w:3}})

db.new_collection.find().readConcern("majority")
```