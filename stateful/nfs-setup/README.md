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

## use nfs
```
kubectl apply -f busybox-nfs.yaml
# this pod will log to a volume /mnt/log.txt, which is an nfs volume mopunted to /exports

# to test it is working
cat /var/nfs/exports/log.txt

you should see some logs
```

