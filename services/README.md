# services
## manually
- $ kubectl expose rc hello-rc --name=hello.svc --target-port=800 --type=NodePort
- $ kubectl describe svc hello-svc
- $ kubectl get svc
- $ kubectl delete svc hello-
- kubectl describe pods | grep app

### expose deployment
- kubectl get deployments
- kubectl expose deployment rolling-nginx --port=80 --type=NodePort
- kubectl get svc
- kubectl delete svc rolling-nginx


## DNS
- kubectl create -f dns.yml
- kubectl exec -it busybox2  -- nslookup nginx



## first
- kubectl create -f first.yml
- kubectl describe svc
- kubectl get svc

to access service: http://172.17.223.180:30001/
where 172.17.223.180 is your node ip address.



