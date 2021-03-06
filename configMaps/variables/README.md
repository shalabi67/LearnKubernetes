# Using env variables from a file

## Run
### Create config map from a file
- kubectl create cm variables --from-file=variables
- kubectl get cm
- kubectl get cm variables
- kubectl describe cm variables

### used config map in a pod
- kubectl create -f pod.yaml
- kubectl logs cm-test-pod

## clean up
- kubectl delete -f pod.yaml
- kubectl delete cm variables
