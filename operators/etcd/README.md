# ETCD example
This example exists in [kubernetes-operators-book](ttps://github.com/kubernetes-operators-book/chapters/tree/master/ch03)

## CRs: Custom API Endpoints
A CRD defines the types and values within an instance of a CR. This example defines a new kind of resource, 
the EtcdCluster.
- kubectl create -f crd.yaml
- kubectl get crd

## Defining an Operator Service Account
We just want to take a first look at basic declarations for a service account and the capabilities that account 
needs to run the etcd Operator.
- kubectl create -f service-account.yaml
- kubectl get serviceaccounts

### Role
We give the role a name that we’ll use to reference it from other places: etcd-operator-role. 
The YAML goes on to list the kinds of resources the role may use, and what it can do with them, that is, 
what verbs it can say
- kubectl create -f role.yaml
- kubectl get roles

### RoleBinding
RoleBinding, assigns the role to the service account
- kubectl create -f rolebinding.yaml
- kubectl get rolebindings

## Deploying the etcd Operator
The Operator is a custom controller running in a pod, and it watches the EtcdCluster CR you defined earlier. 
The manifest file deployment.yaml lays out the Operator pod’s specification, including the container image for the 
Operator you’re deploying. Notice that it does not define the spec for the etcd cluster. You’ll describe the desired 
etcd cluster to the deployed etcd Operator in a CR once the Operator is running.
- kubectl create -f deployment.yaml
- kubectl get deployments
- kubectl get pods

## Declaring an etcd Cluster
In cry.yaml file, you created a CRD defining a new kind of resource, an EtcdCluster. 
Now that you have an Operator watching EtcdCluster resources, you can declare an EtcdCluster with your desired state. 
To do so, provide the two spec elements the Operator recognizes: size, the number of etcd cluster members, 
and the version of etcd each of those members should run.

This brief manifest declared in cr.yaml, declares a desired state of three cluster members, each running 
version 3.1.10 of the etcd server.

- kubectl create -f cr.yaml --save-config
- kubectl get pods -w

This example etcd cluster is a first-class citizen, an EtcdCluster in your cluster’s API. Since it’s an API resource, 
you can get the etcd cluster spec and status directly from Kubernetes.

- kubectl describe etcdcluster/example-etcd-cluster
- kubectl describe etcdcluster example-etcd-cluster

## Exercising etcd
You now have a running etcd cluster. The etcd Operator creates a Kubernetes service in the etcd cluster’s namespace. 
A service is an endpoint where clients can obtain access to a group of pods, even though the members of the group may 
change. A service by default has a DNS name visible in the cluster. The Operator constructs the name of the service 
used by clients of the etcd API by appending -client to the etcd cluster name defined in the CR. Here, the client 
service is named example-etcd-cluster-client, and it listens on the usual etcd client IP port, 2379.

- kubectl get services --selector etcd_cluster=example-etcd-cluster

You can run the etcd client on the cluster and use it to connect to the client service and interact with the etcd API.
- kubectl run --rm -i --tty etcdctl --image quay.io/coreos/etcd --restart=Never -- /bin/sh

in the shell you can execute:
- export ETCDCTL_API=3
- export ETCDCSVC=http://example-etcd-cluster-client:2379
- etcdctl --endpoints $ETCDCSVC put foo bar
- etcdctl --endpoints $ETCDCSVC get foo
- exit

## Scaling the etcd Cluster
You can grow the etcd cluster by changing the declared size specification. Edit cr.yaml and change size from 3 to 4 
etcd members.

- kubectl apply -f cr.yaml

## Upgrade etcd version
This could be a complicated manual process, but using operator it is easy.

- kubectl describe etcdcluster example-etcd-cluster
- notice Current Version:         3.1.10
- kubectl get pod example-etcd-cluster-92cz7p5w5b -o yaml
- serach for image: 
- notice image: quay.io/coreos/etcd:v3.1.10

Now let us upgrad. 

- Edit etcd-cluster-cr.yaml and change the version spec from 3.1.10 to 3.2.13. Then apply the new spec to the resource on the cluster:
- kubectl apply -f cr.yaml
- kubectl get pods -w
- kubectl describe etcdcluster example-etcd-cluster
- notice Current Version:         3.2.13


## clean up
- kubectl delete -f service-account.yaml
- kubectl delete -f role.yaml
- kubectl delete -f rolebinding.yaml
- kubectl delete -f crd.yaml
- kubectl delete -f deployment.yaml
- kubectl delete -f cr.yaml




