# PromQl
- [general](general.md)
- [rate and sum](rate-sum.md)

## setup environment
- configure dns: 
```
sudo nano /etc/hosts
add:
192.168.172.12  promo
```
  
- start machines to provide metrics
```
docker run -d -p 8281:80 --name batch2 sixeyed/prometheus-demo-batch:linux
docker run -d -p 8280:80 --name batch sixeyed/prometheus-demo-batch:linux

docker run -d -p 8180:80  --name web sixeyed/prometheus-demo-web:linux
docker run -d -p 8181:80  --name web2 sixeyed/prometheus-demo-web:linux
```

- start prometheus
```
cd monitoring/PromQL

docker run -d --add-host=promo:192.168.178.12 --name=grafana -p 3000:3000 grafana/grafana

docker run \
    -p 9090:9090 \
    -v ~/learn/LearnKubernetes/monitoring/PromQL:/etc/prometheus \
    --name=prometheus \
    --add-host=promo:192.168.178.12 \
    prom/prometheus \
    --config.file=/etc/prometheus/prometheus.yml \
    --web.enable-lifecycle 
```
