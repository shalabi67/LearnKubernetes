# Jobs

## Simple Job
- kubectl create -f simple-job.yml
- kubectl get pods
notice that the status for simple-job is completed.
notice also that ready is 0/1 because the job has finished.

- kubectl get jobs
- kubectl get jobs simple-job -o yaml
- kubectl delete jobs simple-job

## Multiple jobs
- kubectl create -f multible-job.yml
- kubectl get pods
eventually you will be able to see 5 multi-job pods.

## Cron job
- kubectl create -f cron.yml
- kubectl get cronjobs
- kubectl get cronjobs.batch
- kubectl get jobs --watch
this wil keep running because of --watch


### get job logs
- kubectl get pods
- kubectl logs hello-cron-job-1586618400-n8wwr

### clean up
- kubectl delete cronjobs hello-cron-job









