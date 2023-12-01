# node-exporter

## build
```
docker build -t shalabi67/exporter:trafic .
docker push shalabi67/exporter:trafic

docker build -t shalabi67/exporter:notrafic .
docker push shalabi67/exporter:notrafic
```
## setup
- download node-exporter
```
wget https://github.com/prometheus/node_exporter/releases/download/v1.3.0/node_exporter-1.3.0.linux-amd64.tar.gz
tar xvfz node_exporter-1.3.0.linux-amd64.tar.gz
cd node_exporter-1.3.0.linux-amd64.tar.gz

sudo lsof -i -P -n | grep  9100
sudo kill -9 1567
```

- setup node-exporter as a service
```
sudo useradd -M -r -s /bin/false node_exporter
sudo cp node_exporter /usr/local/bin/
sudo chown node_exporter:node_exporter /usr/local/bin/node_exporter
sudo nano  /etc/systemd/system/node_exporter.service
[Unit]
Description=Prometheus Node Exporter
Wants=network-online.target
After=network-online.target
[Service]
User=node_exporter
Group=node_exporter
Type=simple
ExecStart=/usr/local/bin/node_exporter
[Install]
WantedBy=multi-user.target

sudo systemctl daemon-reload
sudo systemctl start node_exporter
sudo systemctl enable node_exporter

curl localhost:9100/metrics
```