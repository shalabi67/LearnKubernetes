# Install prometheus using docker
 
## install
```
cd monitoring/docker/static-configuration/

docker run \
    -p 9090:9090 \
    -v ~/learn/LearnKubernetes/monitoring/docker/static-configuration:/etc/prometheus \
    prom/prometheus
  
```

## test installation
http://localhost:9090/
http://localhost:9090/metrics
http://localhost:9090/classic/config
http://localhost:9090/service-discovery#prometheus
http://localhost:9090/graph?g0.expr=go_goroutines&g0.tab=1&g0.stacked=0&g0.show_exemplars=0&g0.range_input=1h
http://localhost:9090/graph?g0.expr=prometheus_http_requests_total&g0.tab=1&g0.stacked=0&g0.show_exemplars=0&g0.range_input=1h

