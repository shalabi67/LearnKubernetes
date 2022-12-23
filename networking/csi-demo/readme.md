# CSI

## deploy
```
kubectl apply -f ./csi-demo/Deployment.yaml
```

## Demo1
In this demo we will look at the od ips and node ips
```
kubectl get nodes -o wide
```

Now if we lok at the pod ips we notice that they will start with *192.168* which is based on the *podCIDR* we specify in 
the kind configuration.
```
kubectl get pods -o wide
```