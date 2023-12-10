# Initialize machines
This will specify static ips, host names and install containerd.

|  Machine name  |  Ip address  |          Connection           |
|:--------------:|:------------:|:-----------------------------:|
| Control-plane  | 192.168.3.10 | sh -p 2222 mohammad@localhost |
|    Worker1     | 192.168.3.20 | sh -p 2223 mohammad@localhost |
|    Worker2     | 192.168.3.30 | sh -p 2223 mohammad@localhost |


## initialize control plane

#### configure static ip
```
sudo su
nano /etc/netplan/00-installer-config.yaml

#add these
    enp0s8:
      dhcp4: false
      dhcp6: false
      addresses:
        - 10.10.1.10/24
```

```
netplan apply
```

#### configure control-plane
```
bash <<EOF2
sed -i 's/ubuntu/control-plane/g' /etc/hostname
sed -i 's/ubuntu/control-plane/g' /etc/hosts
hostname pc-1
cat >> /etc/hosts <<EOF

10.10.3.10 control-plane
10.10.3.20 worker1
10.10.3.30 worker2
EOF
EOF2
reboot
```

### pre-install containerd
1. Create the configuration file for containerd:
```
cat <<EOF | sudo tee /etc/modules-load.d/containerd.conf
overlay
br_netfilter
EOF
modprobe overlay
modprobe br_netfilter
```

2. Set the system configurations for Kubernetes networking:
```
cat <<EOF | sudo tee /etc/sysctl.d/99-kubernetes-cri.conf
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-ip6tables = 1
EOF
sysctl --system
```

### Install containerd
```
apt install curl gnupg2 software-properties-common apt-transport-https ca-certificates -y
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
add-apt-repository -y "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
apt-get update && apt-get install -y containerd.io
```

### configure containerd
```
mkdir -p /etc/containerd
containerd config default | sudo tee /etc/containerd/config.toml
systemctl restart containerd
```

### disable swap
```
swapoff -a
sed -e '/swap/ s/^#*/#/' -i /etc/fstab
systemctl mask swap.target 
```

### install kublet
```
apt-get update && apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update
apt-get install -y kubelet=1.27.0-00 kubeadm=1.27.0-00 kubectl=1.27.0-00
apt-mark hold kubelet kubeadm kubectl

```

## initialize worker1

#### configure static ip
```
sudo su
nano /etc/netplan/00-installer-config.yaml

#add these
    enp0s8:
      dhcp4: false
      dhcp6: false
      addresses:
        - 10.10.1.20/24
```

```
netplan apply
```

#### configure control-plane
```
bash <<EOF2
sed -i 's/ubuntu/worker1/g' /etc/hostname
sed -i 's/ubuntu/worker1/g' /etc/hosts
hostname pc-1
cat >> /etc/hosts <<EOF

10.10.3.10 control-plane
10.10.3.20 worker1
10.10.3.30 worker2
EOF
EOF2
reboot
```

### pre-install containerd
1. Create the configuration file for containerd:
```
cat <<EOF | sudo tee /etc/modules-load.d/containerd.conf
overlay
br_netfilter
EOF
modprobe overlay
modprobe br_netfilter
```

2. Set the system configurations for Kubernetes networking:
```
cat <<EOF | sudo tee /etc/sysctl.d/99-kubernetes-cri.conf
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-ip6tables = 1
EOF
sudo sysctl --system
```

### Install containerd
```
apt install curl gnupg2 software-properties-common apt-transport-https ca-certificates -y
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
add-apt-repository -y "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
apt-get update && apt-get install -y containerd.io
```

### configure containerd
```
mkdir -p /etc/containerd
containerd config default | sudo tee /etc/containerd/config.toml
systemctl restart containerd
swapoff -a
```

### install kublet
```
apt-get update && apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update
apt-get install -y kubelet=1.27.0-00 kubeadm=1.27.0-00 kubectl=1.27.0-00
apt-mark hold kubelet kubeadm kubectl

```

## initialize worker2

#### configure static ip
```
sudo su
nano /etc/netplan/00-installer-config.yaml

#add these
    enp0s8:
      dhcp4: false
      dhcp6: false
      addresses:
        - 10.10.1.30/24
```

```
netplan apply
```

#### configure hostnames
```
bash <<EOF2
sed -i 's/ubuntu/worker2/g' /etc/hostname
sed -i 's/ubuntu/worker2/g' /etc/hosts
hostname pc-1
cat >> /etc/hosts <<EOF

10.10.3.10 control-plane
10.10.3.20 worker1
10.10.3.30 worker2
EOF
EOF2
reboot
```

### pre-install containerd
1. Create the configuration file for containerd:
```
cat <<EOF | sudo tee /etc/modules-load.d/containerd.conf
overlay
br_netfilter
EOF
modprobe overlay
modprobe br_netfilter
```

2. Set the system configurations for Kubernetes networking:
```
cat <<EOF | sudo tee /etc/sysctl.d/99-kubernetes-cri.conf
net.bridge.bridge-nf-call-iptables = 1
net.ipv4.ip_forward = 1
net.bridge.bridge-nf-call-ip6tables = 1
EOF
sudo sysctl --system
```

### Install containerd
```
apt install curl gnupg2 software-properties-common apt-transport-https ca-certificates -y
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | apt-key add -
add-apt-repository -y "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"
apt-get update && apt-get install -y containerd.io
```

### configure containerd
```
mkdir -p /etc/containerd
containerd config default | sudo tee /etc/containerd/config.toml
systemctl restart containerd
```

### disable swap
```
swapoff -a
sed -e '/swap/ s/^#*/#/' -i /etc/fstab
systemctl mask swap.target 
```

### install kublet
```
apt-get update && apt-get install -y apt-transport-https curl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF | sudo tee /etc/apt/sources.list.d/kubernetes.list
deb https://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update
apt-get install -y kubelet=1.27.0-00 kubeadm=1.27.0-00 kubectl=1.27.0-00
apt-mark hold kubelet kubeadm kubectl

```

