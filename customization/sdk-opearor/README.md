# SDK operator

## docker
- docker pull ibmgaragecloud/operator-sdk
- if container does not exist
    - docker run  --name opeartor-sdk -it ibmgaragecloud/operator-sdk bash
    - docker cp  C:\Users\mohammad\.kube opeartor-sdk:/root/.kube
- if container exist
    - docker start opeartor-sdk
    - docker exec -it opeartor-sdk  bash
- mkdir first
- cd first
- go mod init first
- operator-sdk init --project-name first
- operator-sdk create api --group ship --version v1beta1 --kind Frigate
- make install


operator-sdk create first-op --api-version=example.com/v1 --kind=Frigate --type=helm


operator-sdk new $OPERATOR_NAME --api-version=example.com/v1 --kind=VisitorsApp --type=helm
