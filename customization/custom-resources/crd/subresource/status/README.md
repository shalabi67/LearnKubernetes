# Status subresource



## Run
- kubectl apply -f crd.yaml
- kubectl apply -f cr.yaml
- kubectl get ssub
- kubectl get ssub status-subresource -o jsonpath='{.status.phase}'
  - notice that you will not see any status
- To change the status
    - kubectl proxy --port 8080
    - open new terminal
    - kubectl get ssub -v=7
    - curl -X GET http://localhost:8080/apis/subresource.com/v1/namespaces/default/statuses/status-subresource/status