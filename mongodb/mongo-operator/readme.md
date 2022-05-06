# mongo operator
Setup steps
- create an access key and put it in the 
https://www.youtube.com/watch?v=2Xszdg-4T6A

https://www.mongodb.com/docs/kubernetes-operator/stable/tutorial/install-k8s-operator/


## setup using helm
```
cd mongodb/mongo-operator/
export KUBECONFIG=~/.kube/kind
kind delete cluster --name ops 
kind create cluster --name ops --config kind/kind-config.yaml
kubectl config use-context kind-ops

kubectl label nodes ops-worker2 no-ops-manager=no-ops-manager1
kubectl label nodes ops-worker3 no-ops-manager=no-ops-manager2
kubectl label nodes ops-worker4 no-ops-manager=no-ops-manager3

helm repo add mongodb https://mongodb.github.io/helm-charts
helm install enterprise-operator mongodb/enterprise-operator --namespace mongodb-operator --create-namespace --set operator.watchNamespace=*
helm template mongodb/enterprise-operator \
  -n default \
  --show-only templates/database-roles.yaml | kubectl apply -f -

kubectl describe deployments mongodb-enterprise-operator
```


## ops manager
when creating the secret remember to get the mongo-operator pod ip address
```
kubectl apply -f OpsManager/pods.yaml
kubectl apply -f OpsManager/admin-secret.yaml
kubectl apply -f OpsManager/ops-manager.yaml
kubectl apply -f OpsManager/service-node.yaml
```

## install project
docker exec -it ops-control-plane sh
