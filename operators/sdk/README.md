# SDK-operator installation
we will use a docker image to do the installation.

## Using docker
-from terminal one
    - docker network create --driver=bridge --gateway=192.168.178.1 --subnet=192.168.178.0/25 home
    - docker pull ibmgaragecloud/operator-sdk:0.10.4
    - docker ps
    - docker run --name operator-sdk --network host  -it ibmgaragecloud/operator-sdk:0.10.4 bash
    - docker run --name operator-sdk  -it ibmgaragecloud/operator-sdk:0.10.4 bash
    - GO TO TERMINAL 2
    - cp -r .kube/ ~/.kube
    - echo "192.168.178.10   kubernetes.docker.internal">>/etc/hosts

- from another terminal 2
    - docker cp c:\Users\mohammad\.kube operator-sdk:.kube
    - docker network connect  host operator-sdk

## build operator-sdk
- docker build -t operator-sdk:1.5 .
- docker run --name operator-sdk -it operator-sdk:1.5 bash


