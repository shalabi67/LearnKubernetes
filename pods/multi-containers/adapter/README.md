# Adapter container

## Main app container (web-ctr)
The container image on docker hub [nigelpoulton/nginxadapter:1.0](https://hub.docker.com/r/nigelpoulton/nginxadapter).
The container exposes /nginx_status api, which exposes nginx basic status matrix.

## Adapter container (transformer)
This container consumes /nginx_status api and transfer it to a metrics that can be consumed by prometheus. 

## Run
- kubectl create -f adapter.yaml
- kubectl exec web -it -- bash
- curl localhost/nginx_status
- curl localhost/nginx_status
- on transformer container (adapter container)
- curl localhost:9113/metrics

