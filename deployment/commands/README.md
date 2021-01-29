## Deployment

## example
### Create deployment
To create a deployment you can use the create or apply commands.
- kubectl create -f example.yml --save-config
- kubectl create -f example-svc.yml
- http://172.25.171.154:30002 notice the 30002 is the port name defined in the example-svc yml.
- kubectl describe deployment commands-deploy
- kubectl get rs
- kubectl get pods --show-labels
- kubectl describe ep hello-svc
- kubectl get ep

### Create deployment using apply command
The apply command can be used to create or modify resources.
- kubectl apply -f example.yml
- kubectl get pods
- kubectl scale deployment commands-deploy --replicas=10
- kubectl get pods
- kubectl scale deployment commands-deploy --replicas=3
- kubectl get pods

### Scale deployment


### access deployment
kubectl create -f ./services/first.yml
to access service http://192.168.255.139:30001/
you get the ip either by using hyber-v manager or by getting node InternalIP by running 
- kubectl describe nodes
then look for InternalIP:  192.168.255.139

### update deployment
- kubectl apply -f first-updated.yml --record
- kubectl rollout status deployment hello-deploy
- kubectl get deployment hello-deploy
- kubectl rollout history deployment hello-deploy
- kubectl get rs
- http://192.168.255.139:30001/ to see the new deployment


### rollback to revision 1
- kubectl rollout undo deployment hello-deploy --to-revision=1
- kubectl get deployment


 



