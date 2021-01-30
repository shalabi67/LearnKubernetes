# Storage Class example

## install
- md /c/temp/data/db on windows or /temp/data/db on linux
- if using windows desk top add path /c/temp/data/db to your docker using the setting/resources/file sharing.
- kubectl create --save-config -f config-map.yaml
- kubectl create --save-config -f storage-class.yaml
- kubectl create --save-config -f pv.yaml
- kubectl create --save-config -f pvc.yaml
- kubectl create --save-config -f service.yaml
- kubectl create --save-config -f stateful-set.yaml

## using mongo
- kubectl get pods
- kubectl exec {pod name} -it sh
- mongo

