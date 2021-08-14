# Exercise To test knowledge

## test 1
Configure a node in such a way that no new pods will be scheduled on it, but existing pods will not be moved away.

### solution 1
```
kubectl taint node example-pool-62shrzux3p nopods=test1:NoSchedule
kubectl apply -f normal.yaml
kubectl get pods -o wide
```
You will notice no pods are scheduled on node example-pool-62shrzux3p

### solution 2
A butter solution is to use cordon

```
let us clean up
kubectl taint node example-pool-62shrzux3p nopods=test1:NoSchedule-
kubectl delete -f normal.yaml

do the task

kubectl cordon example-pool-62shrzux3p
kubectl get nodes -o wide
```
you will find the node status is Ready,`SchedulingDisabled`

```
kubectl apply -f normal.yaml
kubectl get pods -o wide
```

## test 2
mark a node with the node label disk=ssd and start a new nginx pod that uses nodeAffinity to be scheduled 
on nodes that have the label disk=SSD set.
```
kubectl label nodes example-pool-c3jkpxppw3 disk=SSD
kubectl apply -f node-selector-disk-ssd.yaml
kubectl get pods -o wide
```
you will notice that pod is running on node example-pool-c3jkpxppw3

## clean up
- kubectl uncordon example-pool-62shrzux3p
- kubectl delete -f normal.yaml
- kubectl delete -f node-selector-disk-ssd.yaml
- 
