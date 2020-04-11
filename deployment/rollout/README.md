# rollout example

## example
- kubectl create -f example.yml
- kubectl get deployments
- kubectl rollout history deployment rolling-nginx

### update deployment
- kubectl edit deployment rolling-nginx
and change image to nginx:1.15 then save changes.

- kubectl rollout history deployment rolling-nginx
- kubectl rollout history deployment rolling-nginx --revision=1
- kubectl rollout history deployment rolling-nginx --revision=2
- kubectl get rs

### roll back to revision 1
- kubectl rollout undo deployment rolling-nginx --to-revision=1
- kubectl rollout history deployment rolling-nginx
- kubectl describe deployments rolling-nginx



