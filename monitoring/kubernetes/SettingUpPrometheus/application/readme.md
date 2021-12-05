# install application and collect its metrics
This example shows how to use nodejs to collect a metrics and deploy the app is a pod.

notice how we configure pod
```
   annotations:
        prometheus.io/scrape: "true"
        prometheus.io/path: 'swagger-stats/metrics'
        prometheus.io/port: "3000"
```

## setup
```
# pasword is normal
docker login -u shalabi67 
docker build -t shalabi67/comicbox application/
docker push shalabi67/comicbox
kubectl apply -f application/deployment.yml

```