# thermometer
This is had benn taken from [Kubernetes Extensibility: Custom Resource Definitions](https://learning.oreilly.com/scenarios/kubernetes-extensibility-custom/9781492083849/)

## run
- kubectl apply -f trm-crd.yaml
- kubectl get crds

### REST management
- kubectl api-resources --api-group='d2iq.com' -o wide
- kubectl get -v=9 --raw /apis/d2iq.com/v1beta1/thermometers/ | jq

### create cr
- kubectl apply -f trm.yaml
- kubectl get trms
- kubectl get trms -A  # get cross all name spaces

### create cr with spec
This example will create a custom resource with spec information that does not exist on crd. 
- kubectl create namespace sweden
- kubectl apply -f trm-invalid.yaml --validate=false
- kubectl get trm stockholm -n sweden -o json | jq  # we will not see the extra data.

### update crd to allow additional data
- kubectl apply -f trm-crd-with.yaml
- kubectl get trm -A  # now we will see the additional data but without any values
- kubectl apply -f crm-valid.yaml
- kubectl get trm -A
