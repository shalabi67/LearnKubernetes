# DevOps Monitoring Deep Dive
[course](https://learn.acloud.guru/course/97037e05-88ed-41a1-92ee-f5a8080318c2/dashboard)

## kind cluster
```
kind create cluster --config setup/kind-config.yaml
```
## setup prometheus
```
kubectl apply -f prometheus/namespaces.yml

# get ip adress  NO NEED TO DO THAT
docker ps
docker inspect kind-worker
docker inspect kind-control-plane

# change ip address in prometheus-config-map.yml file by replacing <KUBERNETES_IP> with the host ip address
kubectl apply -f prometheus/prometheus-config-map.yml

kubectl apply -f prometheus/prometheus-deployment.yml

kubectl apply -f prometheus/prometheus-service.yml

kubectl apply -f prometheus/clusterRole.yml

kubectl apply -f prometheus/kube-state-metrics.yml

http://localhost:5050/
```

## setting grafana
```
kubectl apply -f grafana/grafana-deployment.yml

kubectl apply -f grafana/grafana-service.yml

http://localhost:8000
usename: admin
password: password
```
## setup app
- build application
```
docker build -t shalabi67/devopsmon-app devopsmon-app/
docker login -u shalabi67
docker push shalabi67/devopsmon-app
```

- deploy application
```
kubectl apply -f devopsmon-app/deployment.yml

http://localhost:8001/
```

## setup exporter with traffic
```
kubectl apply -f exporter/deployment.yml
kubectl apply -f exporter/service.yml
localhost:9200/metrics
localhost:9100/metrics
```