## Deployment

## first example
### create deployment
- kubectl create -f first.yml
- kubectl describe deployment hello-deploy
- kubectl get rs

### access deploymnet
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


 



