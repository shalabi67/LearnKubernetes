# Canary Deployment
This lab show you how to do canary deployment. 
The lab uses an image that has two versions 1 and 2. version 2 of the image has canary text on it.

## Run lab
Notice that this sample uses two service one as PortNode and another as load balancer, just to show how this will work on both.
you should use just one service.
- kubectl create -f stable.yml --save-config --record
- kubectl create -f service.yml 
- kubectl create -f svc-lb.yml
- http://172.25.171.154:30011/
- minikube tunne
- kubectl get svc
- http://10.101.0.83/
- kubectl create -f canary.yml --save-config --record

Now, it the canry version is working ok, you can start increasing canary version replicas and decrease stable version replicas.
