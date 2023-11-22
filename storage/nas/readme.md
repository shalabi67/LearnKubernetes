# connect to nas

## setup
```
sudo apt install nfs-common
sudo mkdir /mnt/kubernetes
sudo mount -t nfs -o user=admin 192.168.178.30:/nfs/kubernetes /mnt/kubernetes
```

## persistence volume
```
kubectl apply -f nas/pv.yaml
```

## pvc
```
kubectl apply -f nas/pvc.yaml
```

## deployment
```
kubectl apply -f nas/deployment.yaml
kubetl get pods
kubectl logs PODNAME
```

## cleanup
```
kubectl delete -f nas/deployment.yaml
kubectl delete -f nas/pvc.yaml
kubectl delete -f nas/pv.yaml
```