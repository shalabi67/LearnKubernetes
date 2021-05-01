# Authorization

## run
### create namespace and service account
- kubectl apply -f ns.yaml
- kubectl apply -f serviceaccount.yaml
- kubectl get --namespace auth serviceaccounts    `you should see auth-sa

### optional create alias to stop typing --namespace option
```
 alias kubeauth='kubectl --namespace=auth'
 alias
 
 kubeauth get serviceaccounts  
```

### create role
- kubectl --namespace=auth apply -f role.yaml
- kubectl --namespace=auth get roles  `should see the auth-view role`
- kubectl --namespace=auth describe role auth-view

### create role binding
- kubectl apply -f role-binding.yaml
- kubectl --namespace=auth describe rolebinding auth-binding

### test results
- kubectl auth can-i list pods --as=system:serviceaccount:auth:auth-sa --namespace=auth
- should see yes
- kubectl auth can-i get pods --as=system:serviceaccount:auth:auth-sa --namespace=auth
- should see yes
- kubectl auth can-i list services --as=system:serviceaccount:auth:auth-sa --namespace=auth
- should see no

