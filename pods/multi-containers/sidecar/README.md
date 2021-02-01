# Sidecar containers
In this example we have two app containers. since they are containers they will run in parallel.
One of the containers will sync data from git to a shared volume. this container will keep syncing from git.
The other container is a web app conatiner, and will show the data the other container synced from git.

## run
- kubectl create -f sidecar.yaml
- kubectl get svc
- http://localhost/
- kubectl exec git-syncer -c web -it -- bash
- ls /usr/share/nginx 
- cat /usr/share/nginx/html/index.html

