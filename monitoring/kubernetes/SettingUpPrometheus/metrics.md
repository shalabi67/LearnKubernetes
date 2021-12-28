# getting metrics

## nodes
```
kubectl exec -it prometheus-deployment-7596b86cb9-79872 -n monitoring -- sh
export TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
wget --no-check-certificate  --header="Authorization: Bearer $TOKEN" https://kubernetes.default.svc/api/v1/nodes/kind-control-plane/proxy/metrics
mv metrics metrics-kind-control-plane
wget --no-check-certificate  --header="Authorization: Bearer $TOKEN" https://kubernetes.default.svc/api/v1/nodes/kind-worker/proxy/metrics
mv metrics metrics-kind-worker

```

## cadvicer
```
kubectl exec -it prometheus-deployment-7596b86cb9-79872 -n monitoring -- sh
export TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
wget --no-check-certificate  --header="Authorization: Bearer $TOKEN" https://kubernetes.default.svc/api/v1/nodes/kind-worker/proxy/metrics/cadvisor
mv cadvisor cadvisor-worker
wget --no-check-certificate  --header="Authorization: Bearer $TOKEN" https://kubernetes.default.svc/api/v1/nodes/kind-control-plane/proxy/metrics/cadvisor
mv cadvisor cadvisor-control

```

## api server
```
kubectl exec -it prometheus-deployment-7596b86cb9-79872 -n monitoring -- sh
export TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
wget --no-check-certificate  --header="Authorization: Bearer $TOKEN" https://172.20.0.2:6443/metrics
mv metrics metrics-api-server

```

## end points
```
kubectl exec -it prometheus-deployment-7596b86cb9-79872 -n monitoring -- sh
wget http://10.244.1.2:9153/metrics
mv metrics metrics-1-2

```