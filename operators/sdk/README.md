# SDK-operator installation
we will use a docker image to do the installation.

## Using docker
-start operator-sdk
    - docker pull shalabi67/operator-sdk:1.5.2
    - docker run --name operator-sdk  -it shalabi67/operator-sdk:1.5.2 bash

    - if container exist
        - docker start operator-sdk
        - docker exec -it operator-sdk  bash

- from another terminal 2
    - docker cp c:\Users\mohammad\.kube operator-sdk:.kube
    - docker cp /home/mohammad/.kube operator-sdk:.kube
    - docker network connect  host operator-sdk

- run on operator-sdk container
    - cp -r .kube/ ~/.kube
    - echo "192.168.178.11   kubernetes.docker.internal">>/etc/hosts
    - echo "export DOCKER_HOST=tcp://192.168.178.11:2375" >> ~/.bashrc && source ~/.bashrc

## build operator-sdk
- docker build -t shalabi67/operator-sdk:1.5 .
- docker run --name operator-sdk -it shalabi67/operator-sdk:1.5 bash


# Not used
    - docker network create --driver=bridge --gateway=192.168.178.1 --subnet=192.168.178.0/25 home
    - docker pull shalabi67/operator-sdk:1.5.2
    - docker ps
    - # docker run --name operator-sdk --network host  -it ibmgaragecloud/operator-sdk:0.10.4 bash
    - docker run --name operator-sdk  -it shalabi67/operator-sdk:1.5.2 bash