## Using env variables with docker
### ARG using env variable
- Build image: `docker build -t env-arg --build-arg APP=APP1 --build-arg DEPLOY=12 .`
- docker run --env STAGE=live -it --rm env-arg 
- docker exec -it test /bin/bash