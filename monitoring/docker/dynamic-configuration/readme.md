# Prometheus dynamic configurations

## Install node exporter
- download node exporter from https://prometheus.io/download/#node_exporter and get the latest download file link
- open new terminal
```
wget https://github.com/prometheus/node_exporter/releases/download/v1.3.0/node_exporter-1.3.0.linux-amd64.tar.gz
tar xvfz node_exporter-1.3.0.linux-amd64.tar.gz
cd node_exporter-1.3.0.linux-amd64
./node_exporter
```

## run prometheus
```
cd monitoring/docker/dynamic-configuration/

docker run -d --add-host=promo:192.168.178.12 --name=grafana -p 3000:3000 grafana/grafana

docker run \
    -p 9090:9090 \
    -v ~/learn/LearnKubernetes/monitoring/docker/dynamic-configuration:/etc/prometheus \
    --name=prometheus \
    --add-host=promo:192.168.178.12 \
    prom/prometheus \
    --config.file=/etc/prometheus/prometheus.yml \
    --web.enable-lifecycle   
```
