# Labels

## manually adding label
- kubectl get deployments
- kubectl label deployment redis example=adding-label
- kubectl get deployments --show-labels
- 

- kubectl get pods
- kubectl label pods frontend app=frontend
- kubectl get pods --show-labels
- kubectl get pods --selector=app=frontend

### remove label
- kubectl label pods frontend app-
- kubectl get pods --show-labels

