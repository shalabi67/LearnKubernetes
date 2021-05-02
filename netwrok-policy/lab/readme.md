# Lab

## pre-setup
- install [calico](https://docs.projectcalico.org/getting-started/kubernetes/quickstart)
`kubectl create -f canal.yaml`

  
Notice that you may need to change cider

- verify
`watch kubectl get pods -n calico-system`
  
- Remove the taints on the master so that you can schedule pods on it.
`kubectl taint nodes --all node-role.kubernetes.io/master-`
