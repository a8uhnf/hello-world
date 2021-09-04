#!/bin/bash

echo "building docker image with gcloud, docker"

tag=$(cat tag.txt)

echo $tag

docker build . -t a8uhnf/gcloud-docker:$tag

docker push a8uhnf/gcloud-docker:$tag