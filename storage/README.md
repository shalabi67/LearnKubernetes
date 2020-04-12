# Storage

## Volume
in this example we will create two container and share a volume between them. 
then we create a file in one container and show that the other container can access the file.

- kubectl create -f volume.yml
- kubectl get pods volume-example
- kubectl describe pod volume-example
have a look at the container mount section and the volumes section.

- kubectl exec -it volume-example -c centos1 -- touch /centos1/test.dat
this command will create file test.dat on the volume

- kubectl exec -it volume-example -c centos2 -- ls centos2
this will show that centos2 container can show the file test.dat

## Persistent Volume
you can get a documentation on this https://kubernetes.io/docs/tasks/configure-pod-container/configure-persistent-volume-storage/

### create volume 
- minikube ssh
- sudo mkdir /mydata
- sudo sh -c "echo 'Hello from Kubernetes storage' > /mydata/index.html"

### create Persistent Volume
- kubectl create -f PersistentVolume.yml
- kubectl get pv

### create Persistent Volume Claim
- kubectl create -f PersistentVolumeClaim.yml
- kubectl get pvc

### create Pod
- kubectl create -f pod.yml

### acess pod
- kubectl port-forward pod/pv-pod 8888:80
- http://localhost:8888/
you should see: Hello from Kubernetes storage
