# Setting up prometheus on kubernetes.
- [source code](https://github.com/linuxacademy/content-kubernetes-prometheus-env)
- [Training course: Monitoring Kubernetes With Prometheus](https://learn.acloud.guru/course/97037e05-88ed-41a1-92ee-f5a8080318c2/dashboard)
- [nodejs application source code](https://github.com/linuxacademy/content-kubernetes-prometheus-app)

## setting kubernetes
- from local-setup directory install kind

## setting prometheus
```
kubectl apply -f namespaces.yml

# get ip adress  NO NEED TO DO THAT
docker ps
docker inspect kind-worker
docker inspect kind-control-plane

# change ip address in prometheus-config-map.yml file by replacing <KUBERNETES_IP> with the host ip address
kubectl apply -f prometheus-config-map.yml

kubectl apply -f prometheus-deployment.yml

kubectl apply -f prometheus-service.yml

kubectl apply -f clusterRole.yml

kubectl apply -f kube-state-metrics.yml

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

### setting exporter
```
docker cp ~/Downloads/node_exporter kind-control-plane:/usr/local/bin/
docker cp ~/Downloads/node_exporter kind-worker:/usr/local/bin/

docker exec -it kind-control-plane sh
useradd -M -r -s /bin/false node_exporter
chown node_exporter:node_exporter /usr/local/bin/node_exporter
node_exporter

docker exec -it kind-worker sh
useradd -M -r -s /bin/false node_exporter
chown node_exporter:node_exporter /usr/local/bin/node_exporter
node_exporter
```

## changing configuration
```
kubectl apply -f prometheus-config-map.yml
kubectl get pods -A
kubectl delete pods prometheus-deployment-7596b86cb9-zrpvz -n monitoring
```

## Adding application
kubectl apply -f application/deployment.yml

http://localhost:8001/swagger-stats/metrics
## Adding rules
```
kubectl apply -f readrules/prometheus-config-map.yml
kubectl apply -f readrules/prometheus-read-rules-map.yml
kubectl apply -f readrules/prometheus-deployment.yml

```

## alert manager
- create slack: learnoliver108@gmail.com [slack](https://app.slack.com/client/T02PCMY93N2/C02PGEXBJKF)

```
kubectl apply -f alertmanager/alertmanager-configmap.yml

kubectl apply -f alertmanager/alertmanager-depoloyment.yml

kubectl get pods -A

kubectl apply -f alertmanager/alertmanager-service.yml 

kubectl apply -f alertmanager/prometheus-config-map.yml
 
kubectl apply -f alertmanager/prometheus-rules-config-map.yml

http://localhost:5051/#/alerts 
```