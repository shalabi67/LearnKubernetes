# Schedular
- [Kubernetes docuemntation](https://kubernetes.io/docs/concepts/scheduling-eviction/taint-and-toleration/)
- [Certified Kubernetes Administrator (CKA), 2nd Ed](https://learning.oreilly.com/videos/certified-kubernetes-administrator/9780137438419/9780137438419-CKA2_03_09_03/)
- [source code](https://github.com/sandervanvugt/cka)

## Pods with toleration can schedule on non taint nodes
This example will show that pods with toleration can be schedule on non taint

```
sudo apt  install jq

# get tain in all nodes
kubectl get nodes -o json | jq '.items[].spec.taints'

we should see null on all nodes

kubectl apply -f schedule-on-tainted.yaml
kubectl get pods

we should see all pods are scheduled.

let us clean up
kubectl delete -f schedule-on-tainted.yaml
```

Now let us add taint to nodes

```
Let us taint all nodes
kubectl taint node $(kubectl get nodes | awk '{print $1}') example=1:NoSchedule

# get tain in all nodes
kubectl get nodes -o json | jq '.items[].spec.taints'

kubectl apply -f schedule-on-tainted.yaml

kubectl get pods

we see all pods are scheduled
```

let us clean up
```
kubectl delete -f schedule-on-tainted.yaml
kubectl taint node $(kubectl get nodes | awk '{print $1}') example=1:NoSchedule-

let us verify cleanup
kubectl get nodes -o json | jq '.items[].spec.taints'
kubectl get pods
```
## Pods not schedules in taint nodes
This example will show pods can not be scheduled in taint nodes
```
kubectl get nodes
# taint nodes
kubectl taint node $(kubectl get nodes | awk '{print $1}') example=1:NoSchedule
# ignore error Error from server (NotFound): nodes "NAME" not found

# try to schedule pods, all pods should be pinding
kubectl apply -f no-schedule.yaml

kubectl get pods -o wide
```

Now let us remove the taint from one node. in this case all of our pods will be scheduled in that node.
```
kubectl get nodes
kubectl taint node example-pool-62shrzux3p example=1:NoSchedule-
kubectl get pods -o wide
```

Let us clean up
```
kubectl taint node $(kubectl get nodes | awk '{print $1}') example=1:NoSchedule-
kubectl delete -f no-schedule.yaml
```

## 