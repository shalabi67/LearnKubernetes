## Create and interrogate secrets from the command line with kubectl.
From the Master Node, as cloud_user

Example: Username and Password

First, store the secret data in a file. In this example, we will place a username and password in two files encoded with base64.

`echo -n 'admin' > ./username.txt`
`echo -n 'L1nux@cad3my' > ./password.txt`

The kubectl can package these files into a 'Secret' object on the API server.

`kubectl create secret generic ks-user-pass --from-file=./username.txt --from-file=./password.txt`

You can look up secrets with **get** and **describe** as follows:

`kubectl get secrets`
`kubectl describe secrets/ks-user-pass`

Secrets are masked by default. If you need to obtain the value of a stored secret, you may use the following commands:

`kubectl get secret ks-user-pass -o yaml`

Then decode the values with:

`echo '[stored value here]' | base64 -d`

## Create Secrets using YAML.
You may also create secrets with a YAML file. The following is an example:

Example YAML:

```
apiVersion: v1
kind: Secret
metadata:
    name: ks-lab-secret
type: Opaque
data:
    username: "admin"
    password: "L1nux@cad3my"
```
Additional fields may also be stored in a YAML file.

Use an editor to create `ks-secret-config.yaml`.

`vi ks-secret-config.yaml`
```
apiVersion: v1
kind: Secret
metadata:
    name: ks-secret-config
type: Opaque
stringData:
    config.yaml: |-
        apiUrl: https://ks.api.com/api/v1
        username: admin
        password: L1nux@cad3my
        branchid: branch21
```
Then create the secret with:

`kubectl create -f ks-secret-config.yaml`

You may look at the fields by getting the secret in YAML, and then passing the `config.yaml` field through the decoder.

`kubectl get secret ks-secret-config -o yaml`
`echo '[stored value here]' | base64 -d`

## Pass Secrets to a pod through a mounted volume.
Secrets may be passed to pods through mounted volumes or through environment variables.

The following is an example as to how volumeMounts specified in a pod's YAML file may be used:

`vi ks-pod.yaml`

```
apiVersion: v1
kind: Pod
metadata:
name: ks-pod
namespace: default
spec:
containers:
- name: ks-pod
  image: busybox
  command:
    - sleep
    - "10000"
      volumeMounts:
    - name: ks-path
      mountPath: "/etc/ks-path"
      readOnly: true
      restartPolicy: Never
      volumes:
- name: ks-path
  secret:
  secretName: ks-secret-config
  items:
    - key: config.yaml
      path: config.yaml
      mode: 400

```
Then create the pod.
`kubectl create -f ks-pod.yaml`

After creating the pod, verify it is ready.

`kubectl get pods`

Once the pod is ready, exec a shell in the pod container.

`kubectl exec -it ks-pod -- sh`

Once you are inside the busybox container, lets have a look at our secrets.

```
cd /etc/ks-path
ls -l
cat config.yaml
```

## Pass Secrets to a pod through an environment variable.
Now lets do an example where we can get these secrets through an environment variable.

`vi ks-pod-env.yaml`

```
apiVersion: v1
kind: Pod
metadata:
name: ks-pod-env
spec:
containers:
- name: ks-pod-env
  image: busybox
  command:
    - sleep
    - "10000"
      env:
    - name: SECRET_CONFIG
      valueFrom:
      secretKeyRef:
      name: ks-secret-config
      key: config.yaml
      restartPolicy: Never
```
Now lets create the pod.

`kubectl create -f ks-pod-env.yaml`

Lets go have a look.

`kubectl exec -it ks-pod-env -- sh`

And check our variable.

`echo $SECRET_CONFIG`
