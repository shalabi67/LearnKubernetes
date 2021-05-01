# security context
Shows how security context is used within pod. 
[code](https://github.com/kubernetes/website/tree/master/content/en/examples/pods/security)
[task](https://kubernetes.io/docs/tasks/configure-pod-container/security-context/)


## Run
- kubectl apply -f pod.yaml
- kubectl get pods
- kubectl exec -it security-context-demo -- sh

### run inside pod
- id   you will see: uid=1000 gid=3000 groups=2000
- ps aux   notice the shell is running under user 1000
- cd data/demo/
- echo "hello" >test.txt
- ls -la   notice the file permissions 1000 and 2000
```
id

ps aux

cd data/demo/
echo "hello" >test.txt
ls -la
```

