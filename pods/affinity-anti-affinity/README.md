# Affinity and Anti-Affinity
## kubectl explain deployment.spec.template.spec.affinity.podAffinity.requiredDuringSchedulingIgnoredDuringExecution
     If the affinity requirements specified by this field are not met at
     scheduling time, the pod will not be scheduled onto the node. If the
     affinity requirements specified by this field cease to be met at some point
     during pod execution (e.g. due to a pod label update), the system may or
     may not try to eventually evict the pod from its node. When there are
     multiple elements, the lists of nodes corresponding to each podAffinityTerm
     are intersected, i.e. all terms must be satisfied.

     Defines a set of pods (namely those matching the labelSelector relative to
     the given namespace(s)) that this pod should be co-located (affinity) or
     not co-located (anti-affinity) with, where co-located is defined as running
     on a node whose value of the label with key <topologyKey> matches that of
     any node on which a pod of the set of pods is running

## Anti-Affinity example
This manifest of an NGINX deployment has four replicas and the selector label app=frontend. 
The deployment has a PodAntiAffinity stanza configured that will ensure that the scheduler does not colocate replicas on a single node. 
This ensures that if a node fails, there are still enough replicas of NGINX to serve data from its cache.

### execute
- kubectl create -f anti-affinity-pod.yml
- kubectl get deployments
  NAME                  READY   UP-TO-DATE   AVAILABLE   AGE
  nginx-anti-affinity   1/4     4            1           8s
  
 notice that we have 1 pod of four installed. 
 in this deployment we have one node and based on our Anti-Affinity we can not deploy more than one pod on the same node. thus we have single pod deployed.
 
 ### cleanup
 - kubectl delete deployment nginx-anti-affinity
  

## Affinity example
- kubectl delete deployment nginx-anti-affinity
- kubectl create -f affinity-pod.yml
- kubectl get deployments
  NAME             READY   UP-TO-DATE   AVAILABLE   AGE
  nginx-affinity   4/4     4            4           2m5s
