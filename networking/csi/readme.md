# Networking



## kind setup
![Lab](lab.jpg)
```
kind delete cluster --name networking
kind create cluster --name networking --config=setup/kind-config.yaml
kubectl apply -f https://docs.projectcalico.org/v3.21/manifests/calico.yaml
```

## Demos
* [CSI kaligo demo](csi-calico-demo/readme.md)