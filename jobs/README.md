# Jobs & Cron jobs
This is an example on how to create jobs & cron jobs

## Jobs
- kubectl create -f job.yml --save-config
- kubectl get jobs
- kubectl describe job pie-counter
- kubectl get pods
- kubectl logs pie-counter-d5cqp
- kubectl get job pie-counter -o=wide
- kubectl get job pie-counter -o=json

## Cron Jobs
- kubectl create -f cronjob.yml --save-config
- kubectl get cronjobs
- or  kubectl get cj
- kubectl describe job pie-counter-cron
- kubectl get pods --watch
- kubectl logs pie-counter-cron-1611929460-jztts
- kubectl get cj pie-counter-cron -o=wide
- kubectl get cj pie-counter-cron -o=json

## Cleanup
- kubectl delete job pie-counter
- kubectl delete cj pie-counter-cron
