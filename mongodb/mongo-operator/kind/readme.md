# kind setup
```
kind create cluster --name ops --config kind-config.yaml
```

## install matlab to run load balancer
```
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/namespace.yaml
kubectl apply -f https://raw.githubusercontent.com/metallb/metallb/v0.12.1/manifests/metallb.yaml
```
## setup address range
```
docker network inspect -f '{{.IPAM.Config}}' kind
```

change the address in the metallb-config.yaml based on the above output
```
kubectl apply -f kind/metallb-config.yaml
```
