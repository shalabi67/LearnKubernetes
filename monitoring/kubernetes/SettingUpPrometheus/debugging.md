# debugging

## debugging
### dns lookup
```
kubectl apply -f ../debugging/dns-util.yaml

kubectl get pods dnsutils

kubectl exec -i -t dnsutils -- nslookup kubernetes.default

```

### service map to ip table
- control
```
docker exec -it kind-control-plane sh
iptables -t nat -L -n
```

- worker
```
docker exec -it kind-control-plane sh
iptables -t nat -L -n
```

### get metrics
- worker nodes
```

export TOKEN=$(cat /var/run/secrets/kubernetes.io/serviceaccount/token)
wget --no-check-certificate  --header="Authorization: Bearer $TOKEN" https://kubernetes.default.svc/api/v1/nodes/kind-control-plane/proxy/metrics
mv metrics kind-control-plane
wget --no-check-certificate  --header="Authorization: Bearer $TOKEN" https://kubernetes.default.svc/api/v1/nodes/kind-worker/proxy/metrics


```