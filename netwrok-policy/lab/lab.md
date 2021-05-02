# Network Policy with Calico
## Introduction
This lab covers the Kubernetes feature of Network Policy. The lab utilizes the **kops** installer to create a cluster using the Calico network overlay. The student is guided through the process of first creating a network policy that prohibits pod access, followed by another policy that grants pod access to certain clients and a named server.

## Solution
Log into the server using the credentials provided for the lab.

`ssh cloud_user@<PUBLIC_IP_ADDRESS>`
### Create the Kubernetes Cluster
- The `k8s-create.sh` script should be in the cloud_user's home directory. List the directory contents to verify that the script is there:

`$ ls -l`

- Run the script using the following command:

`$ . ./k8s-create.sh`

Note: Be sure to leave a space between the `.` and the `./` in front of the script. This ensures that environment variables set in the script are then available to the parent shell.

- Once the cluster configuration has been created, apply the configuration with the following command:

`$ kops update cluster -y`

Note: To view the cluster servers as they are being created, you may use the AWS console and credentials provided.

- Validate the cluster using the following command:

`$ kops validate cluster`

Note: It will return errors until the cluster is fully configured. This can take a few minutes to complete.

When complete, it will report that the cluster is ready.

- Verify the cluster is running:

`$ kubectl get nodes`

- View all the pods that are running:

`$ kubectl get pods --all-namespaces`

Note that the kube-system namespace has the Calico nodes running, as well as the kube-controller.

### Configure the Required Namespace
- Create a namespace called `policy-demo`:

`$ kubectl create ns policy-demo`
The command should respond with an affirmation that the namespace was created.

### Create the Demo Pods
- Run two replicas of the nginx service:

`$ kubectl run --namespace=policy-demo nginx --replicas=2 --image=nginx`

- Expose the service on port 80:

`$ kubectl expose --namespace=policy-demo deployment nginx --port=80`

- View the two replica pods that are running nginx:

`$ kubectl get pods --namespace policy-demo`

- Run an interactive session in a pod called `access` using the busybox image:

`$ kubectl run --namespace=policy-demo access --rm -ti --image busybox /bin/sh`

- Once inside the image, verify access to the nginx server:

`/ # wget -q nginx -O -`

This should respond with the raw html from nginx. This response validates that you have network access on port 80 from the pod you executed called `access`, proving that the network policy is open and you have connectivity.

- Exit the interactive container session:

`/ # exit`

### Enable Isolation
- Download the default-deny yaml file:

`$ wget https://raw.github.com/linuxacademy/content-kubernetes-security-ac/master/default-deny.yaml`

- View the contents of the yaml file:

`$ more default-deny.yaml`

Note that `matchlabels` is a null set, `{}`, which means it is going to match on everything. Once you create the policy, all pods will be denied east-west traffic on the network.

- Create the policy:

`$ kubectl create -f default-deny.yaml`

### Test Isolation
- View that the nginx pods are running:

`$ kubectl get pods --namespace=policy-demo`

- Run an interactive container to test access:

`$ kubectl run --namespace=policy-demo access --rm -ti --image busybox /bin/sh`

- Within the access container, enter:

`/ # wget -q --timeout=5 nginx -O -`

You should receive a timeout after 5 seconds.

- Exit the container shell:

`/ # exit`

### Allow Restricted Access Using a Network Policy
- Download the access-nginx yaml file:

`$ wget https://raw.github.com/linuxacademy/content-kubernetes-security-ac/master/access-nginx.yaml`

- View the contents of the yaml file:

`$ more access-nginx.yaml`

Note that there is an `ingress` defined for `matchlabels` of `run:access` and the `podSelector` that the `ingress` runs against is called `nginx`. So, in this case, the nginx service and the access pod will be able to talk to one another under this network policy.

- Create the policy:

`$ kubectl create -f access-nginx.yaml`

### Verify Access to nginx from the access Pod
- Run an interactive pod called access with an interactive shell:

`$ kubectl run --namespace=policy-demo access --rm -ti --image busybox /bin/sh`

- Once inside the container session, test nginx access:

`/ # wget -q --timeout=5 nginx -O -`

This time, the command should not timeout. Since you are inside the access pod, you should be able to access the nginx service under the network policy.

- Exit the container shell:

`/ # exit`

### Verify That Access to nginx Is Not Allowed from Another Pod
- Run a container shell in a pod called not-access:

`$ kubectl run --namespace=policy-demo not-access --rm -ti --image busybox /bin/sh`

- Within the container, attempt to access nginx:

`/ # wget -q --timeout=5 nginx -O -`

This command should timeout after 5 seconds, since the pod named `not-access` is not able to connect to the nginx service under the network policy.

- Exit the container shell:

`/ # exit`

### Delete the Namespace to Clean Up
Delete the namespace and thus terminate the running pods and nullify the network policies created:

`$ kubectl delete ns policy-demo`
