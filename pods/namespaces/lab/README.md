# This lab will demonstrate namespace

## run
- kubectl create -f ns-testing.yml
- kubectl create -f service.yaml
- kubectl get service --namespace=testing
- kubectl create -f deployment.yaml
- kubectl get deployments --namespace=testing
- kubectl get pods --namespace=testing
