# Schedular
- [Kubernetes docuemntation](https://kubernetes.io/docs/concepts/scheduling-eviction/assign-pod-node/)
- [Certified Kubernetes Administrator (CKA), 2nd Ed](https://learning.oreilly.com/videos/certified-kubernetes-administrator/9780137438419/9780137438419-CKA2_03_09_03/)
- [source code](https://github.com/sandervanvugt/cka)
## node selector
We put a label on na node and on the pod we use node selector to schedule the node the pod should run on.
- kubectl get nodes
- kubectl label nodes shalabi-ionos-node-pool-1-53c3kbfk7t disktype=ssd
- kubectl get nodes --show-labels
- kubectl apply -f pod-node-selector.yaml
- kubectl get pods -o wide

## Pod With Node Affinity
- kubectl apply -f pod-with-node-affinity.yaml
- kubectl get pods 
You will notice that the pod is binding and is not scheduled.

- kubectl label nodes shalabi-ionos-node-pool-1-53c3kbfk7t kubernetes.io/e2e-az-name=e2e-az1
- kubectl get pods
You will notice the pod had been scheduled.


## Pod affinity and anti affinity
### Pod affinity
```
kubectl apply -f pod-affinity-example1.yaml
kubectl get pods
```
we will notice the pod is in bending state. To find out why it is in bending state
```
kubectl describe pod pod-affinity-example1
```

Now, let us make a node match our pod affinity
```
kubectl apply -f pod-affinity-security-s1.yaml
kubectl get pods -o wide
```

we notice the pod is running on different node than the security pod is running on.

### Pod anti affinity example 1

```
kubectl apply -f pod-anti-affinity-example1.yaml
kubectl get pods -o wide
```
Now change pod name and keep running this multiple time until you see a binding pod.

## Pod affinity example 2
Here we are specifying that all nginx pods should be scheduled together.

```
kubectl apply -f pod-affinity-example2.yaml
kubectl get pods -o wide
```

### Pod anti affinity example 2
This example will prevent pods with label app=web from running on same node

```
kubectl apply -f pod-anti-affinity-example2.yaml
kubectl get pods -o wide
```

Now change pod name and keep running this multiple time until you see a binding pod.

## Pod anti affinity example 3
This example using deployment to deploy pods to different nodes using pod anti affinity. This example will prevent pods with label app=nginx from running on same node

```
kubectl apply -f pod-anti-affinity-example3.yaml
kubectl get pods -o wide
```

You will notice there are some pods binding. The number of binding pods will depend on the number of replicas and number of nodes on your cluster.

## Example deploy dependent apps
In a three node cluster, a web application has in-memory cache such as redis. We want the web-servers to be co-located with the cache as much as possible.

```
kubectl apply -f pod-affinity-store.yaml
kubectl get pods -o wide
kubectl apply -f pod-affinity-web.yaml
kubectl get pods -o wide
```