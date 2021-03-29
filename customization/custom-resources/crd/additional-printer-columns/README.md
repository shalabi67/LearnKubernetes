# Additional Printer Columns
The kubectl tool relies on server-side output formatting. 
Your cluster's API server decides which columns are shown by the kubectl get command. 
You can customize these columns for a CustomResourceDefinition. 
The following example adds the Spec, Replicas, and Age columns.

<b>Priority</b>

Each column includes a priority field. Currently, the priority differentiates between columns shown in standard view or wide view (using the -o wide flag).

Columns with priority 0 are shown in standard view.
Columns with priority greater than 0 are shown only in wide view


## Run
- kubectl create -f crd.yaml
- kubectl create -f cr.yaml
- kubectl get apcs
- kubectl get apc -o wide
