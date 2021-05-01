# SDK-operator installation
we will use a docker image to do the installation.

This documentation refers to container as the operator-sdk container, and host machine as the machine running container.   

## configure host machine
[configure docker remote access](https://dockerlabs.collabnix.com/beginners/components/daemon/access-daemon-externally.html)
- sudo mkdir -p /etc/systemd/system/docker.service.d
- sudo nano /etc/systemd/system/docker.service.d/options.conf
  [Service]
  ExecStart=
  ExecStart=/usr/bin/dockerd -H unix:// -H tcp://0.0.0.0:2375
  
- sudo systemctl daemon-reload
- sudo systemctl restart docker

## Using docker
- start operator-sdk
    - docker pull shalabi67/operator-sdk:1.5.2
    - docker run --name operator-sdk \
        -v /home/mohammad/learn/LearnKubernetes/operators/projects:/projects \
        -it shalabi67/operator-sdk:1.5.2 bash
    - docker container run -e "DOCKER_HOST=$(ip -4 addr show docker0 | grep -Po 'inet \K[\d.]+')"

    - if container exist
        - docker start operator-sdk
        - docker exec -it operator-sdk  bash

- from another terminal 2
    - docker cp c:\Users\mohammad\.kube operator-sdk:.kube
    - docker cp /home/mohammad/.kube operator-sdk:.kube
    - not used docker network connect  host operator-sdk

- run on operator-sdk container
    - cp -r .kube/ ~/.kube
    - chmod -R 777 /projects/
    - if running windows machine: echo "192.168.178.10   kubernetes.docker.internal">>/etc/hosts
    - echo "export DOCKER_HOST=tcp://172.17.0.1:2375" >> ~/.bashrc && source ~/.bashrc

## build operator-sdk
- docker build -t shalabi67/operator-sdk:1.5 .
- docker run --name operator-sdk -it shalabi67/operator-sdk:1.5 bash

## create an operator
- from container
    - cd projects
    - mkdir {operator}
    - cd {operator}
    - operator-sdk init --domain tutorial.kubebuilder.io --repo tutorial.kubebuilder.io/project
    - operator-sdk create api --group batch --version v1 --kind CronJob --resource --controller
    - operator-sdk create webhook --group batch --version v1 --kind CronJob --defaulting --programmatic-validation    

- from host
  - sudo chmod -R 777 /projects/  # from operator directory run:
  - docker cp operator-sdk:/projects /home/mohammad/learn/LearnKubernetes/operators/projects

## Not used
    - docker network create --driver=bridge --gateway=192.168.178.1 --subnet=192.168.178.0/25 home
    - docker pull shalabi67/operator-sdk:1.5.2
    - docker ps
    - # docker run --name operator-sdk --network host  -it ibmgaragecloud/operator-sdk:0.10.4 bash
    - docker run --name operator-sdk  -it shalabi67/operator-sdk:1.5.2 bash