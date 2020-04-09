# Replication Controller
## first
- kubectl create -f first.yml
- kubectl describe pods
- kubectl get pods
### update
change the number of replicas from 10 to 20 
- kubectl apply -f first.yml
