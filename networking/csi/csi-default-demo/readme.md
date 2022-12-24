# CSI

## deploy
let us create the cluster
```
kind delete cluster --name networking
kind create cluster --name networking --config=setup/kind-default-config.yaml
```

```
kubectl apply -f ./csi-default-demo/Deployment.yaml
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

## Demo2
This demo will show the node routes.

let us start by getting pod ips
```
kubectl get pods -o wide 
NAME                                           READY   STATUS    RESTARTS   AGE     IP            NODE                       NOMINATED NODE   READINESS GATES
hello-world-5f5b76b5bd-cxlfx                   1/1     Running   0          7m38s   192.168.2.2   networking-worker          <none>           <none>
hello-world-5f5b76b5bd-hfhdf                   1/1     Running   0          7m38s   192.168.1.2   networking-worker2         <none>           <none>
hello-world-5f5b76b5bd-hrf4f                   1/1     Running   0          7m38s   192.168.3.2   networking-worker3         <none>           <none>
node-debugger-networking-control-plane-gnb7d   1/1     Running   0          5m17s   172.20.0.3    networking-control-plane   <none>           <none>
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
192.168.0.2     0.0.0.0         255.255.255.255 UH    0      0        0 veth26f7f663
192.168.0.3     0.0.0.0         255.255.255.255 UH    0      0        0 veth6fb2198c
192.168.0.4     0.0.0.0         255.255.255.255 UH    0      0        0 vetha29b7671
192.168.1.0     networking-work 255.255.255.0   UG    0      0        0 eth0
192.168.2.0     networking-work 255.255.255.0   UG    0      0        0 eth0
192.168.3.0     networking-work 255.255.255.0   UG    0      0        0 eth0


# The the node will have these network interface 
ifconfig

eth0: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 172.20.0.3  netmask 255.255.0.0  broadcast 172.20.255.255
        inet6 fe80::42:acff:fe14:3  prefixlen 64  scopeid 0x20<link>
        inet6 fc00:f853:ccd:e793::3  prefixlen 64  scopeid 0x0<global>
        ether 02:42:ac:14:00:03  txqueuelen 0  (Ethernet)
        RX packets 28260  bytes 57793813 (57.7 MB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 18745  bytes 5258449 (5.2 MB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

lo: flags=73<UP,LOOPBACK,RUNNING>  mtu 65536
        inet 127.0.0.1  netmask 255.0.0.0
        inet6 ::1  prefixlen 128  scopeid 0x10<host>
        loop  txqueuelen 1000  (Local Loopback)
        RX packets 84834  bytes 24037076 (24.0 MB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 84834  bytes 24037076 (24.0 MB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

veth26f7f663: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.0.1  netmask 255.255.255.255  broadcast 0.0.0.0
        ether 72:56:21:24:4c:14  txqueuelen 0  (Ethernet)
        RX packets 602  bytes 53103 (53.1 KB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 591  bytes 174093 (174.0 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

veth6fb2198c: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.0.1  netmask 255.255.255.255  broadcast 0.0.0.0
        ether f6:01:4e:0a:ea:f2  txqueuelen 0  (Ethernet)
        RX packets 602  bytes 53128 (53.1 KB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 595  bytes 174344 (174.3 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0

vetha29b7671: flags=4163<UP,BROADCAST,RUNNING,MULTICAST>  mtu 1500
        inet 192.168.0.1  netmask 255.255.255.255  broadcast 0.0.0.0
        ether e6:bc:0a:ee:e2:77  txqueuelen 0  (Ethernet)
        RX packets 1308  bytes 260868 (260.8 KB)
        RX errors 0  dropped 0  overruns 0  frame 0
        TX packets 1300  bytes 421766 (421.7 KB)
        TX errors 0  dropped 0 overruns 0  carrier 0  collisions 0
```

Now let us connect to a node that is running a pod.
```
kubectl get pods -A -o wide | grep networking-worker2

NAMESPACE            NAME                                               READY   STATUS    RESTARTS   AGE     IP            NODE                       NOMINATED NODE   READINESS GATES
default              hello-world-5f5b76b5bd-hfhdf                       1/1     Running   0          9m9s    192.168.1.2   networking-worker2         <none>           <none>
kube-system          kindnet-6zfc2                                      1/1     Running   0          11m     172.20.0.5    networking-worker2         <none>           <none>
kube-system          kube-proxy-xkk64                                   1/1     Running   0          11m     172.20.0.5    networking-worker2         <none>           <none>

kubectl debug node/networking-worker2 -it --image=ubuntu
apt update
apt install net-tools
route

Kernel IP routing table
Destination     Gateway         Genmask         Flags Metric Ref    Use Iface
default         home            0.0.0.0         UG    0      0        0 eth0
172.20.0.0      0.0.0.0         255.255.0.0     U     0      0        0 eth0
192.168.0.0     networking-cont 255.255.255.0   UG    0      0        0 eth0
192.168.1.2     0.0.0.0         255.255.255.255 UH    0      0        0 vethc9baf9d3
192.168.2.0     networking-work 255.255.255.0   UG    0      0        0 eth0
192.168.3.0     networking-work 255.255.255.0   UG    0      0        0 eth0
```