# Using configuration files

## Run
### Create config map from a file
- kubectl create cm nginx-config-cm --from-file=configuration.conf
- kubectl get cm
- kubectl get cm nginx-config-cm
- kubectl describe cm nginx-config-cm

### used config map in a pod
- kubectl create -f pod.yaml
- kubectl exec -it nginx-cm bash
- `ls /etc/nginx/conf.d`
- `cat ls /etc/nginx/conf.d/default.conf`
- `exit`

## clean up
- kubectl delete -f pod.yaml
- kubectl delete cm nginx-config-cm
