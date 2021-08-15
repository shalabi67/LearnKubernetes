# Setup nfs to your local cluster

## docker setup
```
sudo modprobe nfs 

sudo modprobe nfsd 

docker run -d --rm --privileged --name nfs-server  -v /var/nfs:/var/nfs  phico/nfs-server:latest 

docker network list 

# connect kind(my kubernetes cluster) network to nfs server network so they can communicate  

docker network connect kind nfs-server 
```

## use nfs through volumes
```
kubectl apply -f busybox-nfs.yaml
# this pod will log to a volume /mnt/log.txt, which is an nfs volume mopunted to /exports

# to test it is working
cat /var/nfs/exports/log.txt

you should see some logs
```

## use nfs through pvc
```
kubectl apply -f busybox-nfs-pvc-pv.yaml
# this pod will log to a volume /mnt/log.txt, which is an nfs volume mopunted to /exports

# to test it is working
kubectl exec -it busybox-9745544bf-srqjt -- cat /mnt/log.txt
kubectl exec -it busybox-9745544bf-srqjt -- /bin/sh
cat /var/nfs/exports/log.txt

you should see some logs
```

