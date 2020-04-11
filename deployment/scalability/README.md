# scalability

## redis example
- kubectl create -f redis-deploy.yml
- kubectl get deployments
- kubectl get rs
- kubectl get pods

## if we killed pod another one will be created for us
- kubectl delete pods redis-d6c8c987b-k5tr5
- kubectl get pods
now we will see anew pod had been created. that is why we should run using deployments not pods.


