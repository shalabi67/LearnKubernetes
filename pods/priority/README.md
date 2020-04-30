# Priority

## Example
- kubectl create -f priority-class.yml
- kubectl get priorityclass.scheduling.k8s.io
- kubectl create -f pod.yml

### notice
- kubectl describe pod high-priority-pod
...
Priority:             1000
Priority Class Name:  high-priority
...


