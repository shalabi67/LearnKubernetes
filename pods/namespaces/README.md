# Namespaces

## Commands
- kubectl get all --all-namespaces
- kubectl get namespaces

## create namespace
- kubectl create ns production -o yaml
- kubectl delete ns production
- kubectl create -f ns.yml

## use namespace in pod
to use namespace you just put it under metadata as shown in `namespace.yaml`
- kubectl create -f ns-testing.yml
- kubectl create -f pod.yaml

to see you created pod
- kubectl get pods --all-namespaces
- kubectl get pods --namespace=testing

or you can switch the default namespace to testing
- kubectl config set-context --current --namespace=testing
- kubectl get pods