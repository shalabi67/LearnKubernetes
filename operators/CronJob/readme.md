# CronJob
This is a tutorial taken from [CronJob tutorial](https://book.kubebuilder.io/cronjob-tutorial/cronjob-tutorial.html).
This main changed is using operator-sdk framework.

## Step 1: create project
- mkdir cronjob
  cd cronjob
- operator-sdk init --domain learn.com --repo github.com/example/cronjob-operator
- operator-sdk create api --group cronjob --version v1alpha1 --kind CronJob --resource --controller

## Step 2: copy init project
- docker cp operator-sdk:/cronjob/ ./
- add to git
- ignore the following: 
  - zz_generated.deepcopy.go
  

## Step 3: update type
- nano api/v1alpha1/cronjob_types.go
