# create ops manager in kubernets
https://www.mongodb.com/blog/post/running-mongodb-ops-manager-in-kubernetes#:~:text=The%20MongoDB%20Enterprise%20Kubernetes%20Operator%2C%20or%20simply%20the%20Operator%2C%20manages,changing%20these%20settings%20as%20needed.

## setup
```
kubectl apply -f OpsManager/pods.yaml
kubectl apply -f OpsManager/admin-secret.yaml
kubectl apply -f OpsManager/ops-manager.yaml
```