# Secrets

# Example 1
### create secret
- kubectl create secret generic info-secret --from-file=secret1=info.txt --from-literal=passphrase=password1
- kubectl get secrets
- kubectl describe secrets info-secret
- kubectl get secrets info-secret -o yaml


## Example 2: secret from yaml files
### create using base64 
#### encode on linux
- echo -n "lisa" | base64

#### decode
- echo bGlzYQ== | base64 -d

### create using kubernetes
- kubectl create secret generic user-secret --from-literal=user=mohammad --from-literal=passphrase=password1
- kubectl get secrets user-secret -o yaml

copy the data and paste it in your yaml file

### run example
#### create secret
- kubectl create -f user-secret.yml


#### create pod
- kubectl create -f pod.yml
- kubectl get pods secretbox2

#### validate
- kubectl exec -it secretbox2 -- sh
- cat secretstuff/user
mohammad
- cat secretstuff/passphrase
password1/

## Example 3: use secrets as var 

#### create secret
- kubectl create -f user-secret.yml

### create pod
- kubectl create -f pod-secret-var.yml

#### validate
- kubectl exec -it mymysql -- sh
- env
you will find MYSQL_ROOT_PASSWORD=password1

