# Setting up prometheus on kubernetes.
- [source code](https://github.com/linuxacademy/content-kubernetes-prometheus-env)
- [Training course: Monitoring Kubernetes With Prometheus](https://learn.acloud.guru/course/97037e05-88ed-41a1-92ee-f5a8080318c2/dashboard)

## setting kubernetes
- from local-setup directory install kind
- kubectl apply -f https://github.com/coreos/flannel/raw/master/Documentation/kube-flannel.yml

## setting prometheus
```
kubectl apply -f namespaces.yml

docker ps
# get ip adress
docker inspect kind-worker
docker inspect kind-control-plane

# change ip address in prometheus-config-map.yml file by replacing <KUBERNETES_IP> with the ip address
kubectl apply -f prometheus-config-map.yml

kubectl apply -f prometheus-deployment.yml







```
