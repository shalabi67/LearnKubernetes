# CSI

## deploy
let us create the cluster
```
kind delete cluster --name networking
kind create cluster --name networking --config=setup/kind-config.yaml
kubectl apply -f https://docs.projectcalico.org/v3.21/manifests/calico.yaml
```

```
kubectl apply -f ./csi-calico-demo/Deployment.yaml
```

## Demo1
In this demo we will look at the od ips and node ips
```
kubectl get nodes -o wide
```

Now if we look at the pod ips we notice that they will start with *192.168* which is based on the *podCIDR* we specify in 
the kind configuration.
```
kubectl get pods -o wide
```

Now let us have a look at the node network
```
kubectl describe node networking-control-plane
```
look at the annotation and you will see 
```
                    projectcalico.org/IPv4Address: 172.20.0.2/16
                    projectcalico.org/IPv4IPIPTunnelAddr: 192.168.223.128

```
the IPv4Address is the real node address, while IPv4IPIPTunnelAddr is IP address of the tunnel interface that's 
associated with this node. And so the Calico network's going to use tunnels to exchange information between the nodes 
in a cluster so that the Pods running on the nodes in a cluster can use the real IPs independent of the underlying 
infrastructure. Pod traffic will go into the tunnel on one node and pop out of the tunnel on the destination node where
that other Pod is living on that the Pod is trying to communicate to.

move down in the output we will see
```
PodCIDR:                      192.168.0.0/24
PodCIDRs:                     192.168.0.0/24
```
here we have PodCIDR, and then PodCIDR as plural. Now normally what PodCIDR is is the range of IPs that are going to 
be assigned to Pods that are started on this node. Well that's not quite the case inside of the Calico Pod network, 
and we'll look at that in a second, and we'll see that the pods' IPs are going to be allocated from that global 
class pool. Now we're going to compare this with Kubenet, which does use these ranges for allocating IPs from. 
Now you might be wondering why there are two ranges here, PodCIDR and PodCIDRs plural. Well, if it's PodCIDR, 
you're going to be using a single IPv4 or IPv6 range to allocate from, and that's going to be indicated in the 
PodCIDR field. If we look at PodCIDRs plural, if you're using IPv4 and IPv6, then those would be the ranges that it'd 
be allocated from.

## Demo2
This demo will show the node routes.

let us start by getting pod ips
```
kubectl get pods -o wide

And we will see the following
NAME                                           READY   STATUS    RESTARTS   AGE   IP               NODE                       NOMINATED NODE   READINESS GATES
hello-world-7c64bd9f6b-lmhxb                   1/1     Running   0          36m   192.168.21.132   networking-worker3         <none>           <none>
hello-world-7c64bd9f6b-s5cfz                   1/1     Running   0          36m   192.168.78.132   networking-worker          <none>           <none>
hello-world-7c64bd9f6b-tpnrz                   1/1     Running   0          36m   192.168.175.5    networking-worker2         <none>           <none>
node-debugger-networking-control-plane-j8zww   1/1     Running   0          7m    172.20.0.2       networking-control-plane   <none>           <none>
```

Now let us get the route on the control plane node.

```
kubectl debug node/networking-control-plane -it --image=ubuntu
apt update
apt install net-tools
route

The output will be something like this:

Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
default         home            0.0.0.0         UG    0      0        0 eth0
172.20.0.0      0.0.0.0         255.255.0.0     U     0      0        0 eth0
192.168.21.128  networking-work 255.255.255.192 UG    0      0        0 tunl0
192.168.78.128  networking-work 255.255.255.192 UG    0      0        0 tunl0
192.168.175.0   networking-work 255.255.255.192 UG    0      0        0 tunl0
192.168.223.128 0.0.0.0         255.255.255.192 U     0      0        0 *

# The the node will have a network interface called tun0 and we can see that by running
ifconfig
tunl0: flags=193<UP,RUNNING,NOARP>  mtu 1480
        inet 192.168.223.128  netmask 255.255.255.255
        tunnel   txqueuelen 1000  (IPIP Tunnel)
        RX packets 0  bytes 0 (0.0 B)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 0  bytes 0 (0.0 B)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

```

Now let us connect to a node that is running a pod.
```Markdown
kubectl get pods -A -o wide | grep networking-worker2
default              hello-world-7c64bd9f6b-tpnrz                       1/1     Running     0          63m     **192.168.175.5**    networking-worker2         <none>           <none>
kube-system          calico-node-wf5n6                                  1/1     Running     0          124m    172.20.0.5       networking-worker2         <none>           <none>
kube-system          coredns-558bd4d5db-k4j7g                           1/1     Running     0          145m    **192.168.175.1**    networking-worker2         <none>           <none>
kube-system          kube-proxy-8gzgz                                   1/1     Running     0          145m    172.20.0.5       networking-worker2         <none>           <none>

kubectl debug node/networking-worker2 -it --image=ubuntu
apt update
apt install net-tools
route

Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
default         home            0.0.0.0         UG    0      0        0 eth0
172.20.0.0      0.0.0.0         255.255.0.0     U     0      0        0 eth0
192.168.21.128  networking-work 255.255.255.192 UG    0      0        0 tunl0
192.168.78.128  networking-work 255.255.255.192 UG    0      0        0 tunl0
192.168.223.128 networking-cont 255.255.255.192 UG    0      0        0 tunl0
192.168.175.0   0.0.0.0         255.255.255.192 U     0      0        0 *
192.168.175.1   0.0.0.0         255.255.255.255 UH    0      0        0 cali7dac2757ae1
192.168.175.5   0.0.0.0         255.255.255.255 UH    0      0        0 cali52440ec1550
```
Notice that this will have tunnel routes to the other three nodes exiting in the system. 
It also has a route for the two pods running in that node 192.168.175.1 and 192.168.175.5. 
This means that this node can reach all other nodes, and all pods running inside it.

Pod traffic moving between nodes is passed via tunnels,and this traffic goes into tunl0 interfaces based on routes 
defined on the nodes. Once that traffic has reached a destination node for the Pod that it's trying to reach, then 
there's going to be a route on that node that defines an interface that's exposed into the Pod.  