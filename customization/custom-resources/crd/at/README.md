# CRD (CustomResourceDefinition)
CRD is a Kubernetes resource itself. It describes the available CRs in the cluster.

## windows
convert to base64
- certutil -decode fileToConvert.dat base64.dat

## run
- kubectl create -f at-crd.yaml
- kubectl create -f at-cr.yaml
  kubectl get ats
- kubectl get ats -o=wide  This will show <b>command</b> field
- kubectl get ats -v=7


## clean up
- kubectl delete -f at-cr.yaml
- kubectl delete -f at-crd.yaml