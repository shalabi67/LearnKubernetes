# NFS example
This example will show how to connect to my cloud

## create NFS
use my cloude to create a share and configure it to use nfs
nfs://192.168.178.30/nfs/kubernetes

- add index.html to nfs share


## create Persistence volume
- kubectl create -f pv.yml

## create Persistence volume claim
- kubectl create -f pvc.yml
- kubectl get pvc
- kubectl get pv
NAME         CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                   STORAGECLASS   REASON   AGE
mycloud-pv   2Gi        RWX            Retain           Bound    default/mycloud-claim   manual                  10s


## create pod
- kubectl create -f pod.yml

### acess pod
- kubectl port-forward pod/mycloud-pod 8888:80
- http://localhost:8888/
you should see: Hello from mycloud
