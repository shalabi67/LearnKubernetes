# Node selector
allows a pod to run on a specific node based on some labels.

## Example
### Identify Nodes
here we will label minikube node with disktype=ssd
- kubectl get nodes
- kubectl label node minikube disktype=ssd
- kubectl get nodes --show-labels

### run pods
- kubectl create -f pod-hdd.yml
- kubectl create -f pod-ssd.yml

### Validate
notice the status for hdd-pod
- kubectl get pods
NAME                READY   STATUS    RESTARTS   AGE
hdd-pod             0/1     Pending   0          82s
high-priority-pod   1/1     Running   0          57m
ssd-pod             1/1     Running   0          89s

it is clear that the hdd bod has no node to run on, now it is waiting for node with label disktype=ssd

