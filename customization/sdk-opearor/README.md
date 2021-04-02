# SDK operator

## docker
- docker pull ibmgaragecloud/operator-sdk
- if container does not exist
    - docker run  --name operator-sdk -it ibmgaragecloud/operator-sdk bash
    - docker run  --name operator-sdk -it shalabi67/operator-sdk bash
    - docker cp  C:\Users\mohammad\.kube operator-sdk:/root/.kube
    - docker cp  ~/.kube operator-sdk:/root/.kube
- if container exist
    - docker start operator-sdk
    - docker exec -it operator-sdk  bash
- yum -y install nano
- install docker on
  - dnf -y install dnf-plugins-core
  - dnf config-manager \
    --add-repo \
    https://download.docker.com/linux/fedora/docker-ce.repo
  - dnf install docker-ce docker-ce-cli containerd.io
  

  
## using go
- mkdir first
- cd first
- go mod init first
- operator-sdk init --project-name first
- operator-sdk create api --group ship --version v1beta1 --kind Frigate
- make install


## using helm
- operator-sdk init --domain example.com --plugins helm
- operator-sdk create api --group demo --version v1alpha1 --kind Nginx

## Examples
### Helm
This tutoria taken from [tutorial](https://sdk.operatorframework.io/docs/building-operators/helm/tutorial/)
- mkdir nginx-operator
- cd nginx-operator
- operator-sdk init --plugins helm --domain example.com --group demo --version v1alpha1 --kind Nginx

This creates the nginx-operator project specifically for watching the Nginx resource with APIVersion demo.example.com/v1alpha1 and Kind Nginx

- nano helm-charts/nginx/values.yaml
- set replicaCount:2
- set port: 8080
- save and exit
- make install run
- // make docker-build docker-push IMG=quay.io/learnoliver31/nginx-operator:v0.0.1
- docker cp opeartor-sdk:/nginx-operator/ ./
- docker build . -t shalabi67/nginx-operator:v0.0.1
- docker push shalabi67/nginx-operator:v0.0.1
- make deploy IMG=shalabi67E/nginx-operator:v0.0.1


