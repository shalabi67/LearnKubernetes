# Config maps to support environment variables

## Example create manually
- kubectl create cm cm-name --from-literal=VAR=hello --from-literal=VAR1=shalabi
- kubectl create cm cm-name --from-literal=VAR=hello --from-literal=VAR1=shalabi -o yaml --dry-run

## Example
- kubectl create cm variables --from-file=variables
- kubectl get cm
- kubectl describe cm variables

### get yaml file
- kubectl create cm variables1 --from-file=variables -o yaml --dry-run

### create Pod
- kubectl create -f pod.yml
- kubectl logs env-pod


