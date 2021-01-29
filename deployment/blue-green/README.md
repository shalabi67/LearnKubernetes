# Canary Deployment
This lab show you how to do blue-green deployment. 
The lab uses an image that has two versions 3 and 4.

## Service
The service is load balancer service. there are two files of this service one for blue and another for green. 
These two files are same in every thing except role label. 
So, these two files can be replaced by one using environment variable and bash script. 
Please have a look under samples/blue-green in DockerAndKubernetesCourseCode-main.zip

## Run lab 
Notice since services and deployment are same, we will use apply instead of create to switch from blue to green and vise versa

### Deploy green
#### Run on windows
- $Env:TARGET_ROLE="green"
- $Env:IMAGE_VERSION="3.0"
- cat blue-green.yml | %{ $_ -replace "\`$TARGET_ROLE", $Env:TARGET_ROLE } | %{ $_ -replace "\`$IMAGE_VERSION", $Env:IMAGE_VERSION } | kubectl apply --record -f -
- cat service.yml | %{ $_ -replace "\`$TARGET_ROLE", $Env:TARGET_ROLE } | kubectl apply --record -f -

#### run on linux
- export TARGET_ROLE=green
- export IMAGE_VERSION=3.0
- cat blue-green.yml | sh config.sh | kubectl create --save-config -f -
- cat service.yml | sh config.sh | kubectl create --save-config -f -

- http://localhost:8888/


### Deploy blue
#### Run on windows
- $Env:TARGET_ROLE="blue"
- $Env:IMAGE_VERSION="4.0"
- cat blue-green.yml | %{ $_ -replace "\`$TARGET_ROLE", $Env:TARGET_ROLE } | %{ $_ -replace "\`$IMAGE_VERSION", $Env:IMAGE_VERSION } | kubectl apply --record -f -
- cat service.yml | %{ $_ -replace "\`$TARGET_ROLE", $Env:TARGET_ROLE } | kubectl apply --record -f -

#### run on linux
- export TARGET_ROLE=blue
- export IMAGE_VERSION=4.0
- cat blue-green.yml | sh config.sh | kubectl create --save-config -f -
- cat service.yml | sh config.sh | kubectl create --save-config -f -

- http://localhost:8888/