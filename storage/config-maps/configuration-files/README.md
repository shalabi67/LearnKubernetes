# Configuration map as configuration files

## Example
This example shows how to use config maps to provide file configuration. 
The idea is simple use the config map as a volume, then use that volume in your container as shown in the pod.yml file
- kubectl create cm nginy-server-config --from-file serverConfig.txt

### create pod
- kubectl create -f pod.yml
- kubectl exec -it nginx-cm -- sh
- cat /etc/nginx/conf.d/default.conf
