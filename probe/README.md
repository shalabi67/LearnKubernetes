# Probe

## Example
- kubectl create -f pod.yml
- kubectl get pods busybox-ready
NAME            READY   STATUS    RESTARTS   AGE
busybox-ready   0/1     Running   0          82s

what this means is the pod is running but its probe is failing.

- kubectl describe pods busybox-ready
Events:
  Type     Reason     Age                 From               Message
  ----     ------     ----                ----               -------
  Normal   Scheduled  <unknown>           default-scheduler  Successfully assigned default/busybox-ready to minikube
  Normal   Pulling    2m22s               kubelet, minikube  Pulling image "busybox"
  Normal   Pulled     2m20s               kubelet, minikube  Successfully pulled image "busybox"
  Normal   Created    2m20s               kubelet, minikube  Created container busy
  Normal   Started    2m19s               kubelet, minikube  Started container busy
  Warning  Unhealthy  9s (x13 over 2m9s)  kubelet, minikube  Readiness probe failed: cat: can't open '/tmp/nothing': No such file or directory

### fixing
now even the probe is failing but the pod is running, so let us make the probe passing
- kubectl exec -it busybox-ready -- sh
- touch /tmp/nothing

- kubectl get pods busybox-ready
  NAME            READY   STATUS    RESTARTS   AGE
  busybox-ready   1/1     Running   0          6m50s


## Example use tcp socket
- kubectl create -f pod-tcpPort.yml
- kubectl get pods nginx-probes
- kubectl describe pods nginx-probes
....
    Liveness:       tcp-socket :80 delay=20s timeout=1s period=20s #success=1 #failure=3
    Readiness:      tcp-socket :80 delay=5s timeout=1s period=10s #success=1 #failure=3
...
