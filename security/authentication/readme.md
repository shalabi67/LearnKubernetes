# Pod using service account
This example will show hou to use a pod with service account and hot to get service account token

## run
- kubectl apply -f serviceaccount.yaml
- kubectl apply -f pod.yaml

## show token in pod
- kubectl exec -it ks-demo-pod -- sh
- cat /var/run/secrets/kubernetes.io/serviceaccount/token
- https://jwt.io/
- copy token and paste it to jwt website
You will see the following information 
```
{
  "aud": [
    "https://kubernetes.default.svc.cluster.local"
  ],
  "exp": 1651409644,
  "iat": 1619873644,
  "iss": "https://kubernetes.default.svc.cluster.local",
  "kubernetes.io": {
    "namespace": "default",
    "pod": {
      "name": "ks-demo-pod",
      "uid": "518d944d-54ef-4972-b49f-1d5e3812d35f"
    },
    "serviceaccount": {
      "name": "pods-service",
      "uid": "593275b9-d3d0-4f1b-a270-85347e99999e"
    },
    "warnafter": 1619877251
  },
  "nbf": 1619873644,
  "sub": "system:serviceaccount:default:pods-service"
}
```

## cleanup
- kubectl delete -f pod.yaml
-  kubectl delete -f serviceaccount.yaml 
