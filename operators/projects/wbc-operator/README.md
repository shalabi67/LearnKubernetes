# WBC operator
This is an example operator crated using operator sdk. it is following pluralsight training.

## Commands
- operator-sdk init --domain wbc.com --repo learn.oper.com
- operator-sdk create api --group snapshot --version v1alpha1 --kind WbcSnapshot --resource --controller

## Fixing Custom Resource
- update /api/v1alpha1/wbcsnapshot_types.go fields
- make manifests