# init container
this will run before the main container and the main container can not start until the init container finishes.


## wrong
this is an example to show what wil happen if init container did not complete
- kubectl create -f wrong.yml
- kubectl get pods
notice how init-container-will-not-finish pod keeps running in the init status preventing the main container from starting.

### clean up
- kubectl delete pods init-container-will-not-finish

## Valid
this is a valid example
- kubectl create -f valid.yml
- kubectl get pods --watch
