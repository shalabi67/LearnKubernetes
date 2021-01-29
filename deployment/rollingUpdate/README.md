# Rolling update example

## example
- in the example file make sure the replica option is set to 4
- kubectl create -f example.yml

Now, wait until the deployment finishes by running
- kubectl rollout status deployment rolling-nginx

Now, change replicas option from 4 to 20 in the example.yml then run the following commands

- kubectl apply -f example.yml
- kubectl get deployments


- kubectl get pods
- kubectl get deployments
- kubectl rollout history deployment rolling-nginx
