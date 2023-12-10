# initialize cluster
## Control-plane
### setup kubeadm
```
kubeadm init --pod-network-cidr 192.168.0.0/16 --kubernetes-version 1.27.0
```

kubeadm join 10.0.2.15:6443 --token yhuhfi.d3fkmrtz476nlh4l \
--discovery-token-ca-cert-hash sha256:4f497c526eb256f3fd189450b32c00f8d4bafa68631a63197118b03069cea071

### setup kubeconfig
```
mkdir -p $HOME/.kube
cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
chown $(id -u):$(id -g) $HOME/.kube/config
kubectl get nodes
exit
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
kubectl get nodes
```

### Install Calico Network 
```
kubectl apply -f https://raw.githubusercontent.com/projectcalico/calico/v3.25.0/manifests/calico.yaml
```

wait until it installs
```
kubectl get pods -A
kubectl get nodes
```

## worker 1
### join cluster
```
kubeadm join 10.0.2.15:6443 --token yhuhfi.d3fkmrtz476nlh4l \
--discovery-token-ca-cert-hash sha256:4f497c526eb256f3fd189450b32c00f8d4bafa68631a63197118b03069cea071
```

