#!/bin/bash

echo "----"

docker build -t a8uhnf/rabbitmq-cluster:0.0.1 . -f docker/Dockerfile
docker push a8uhnf/rabbitmq-cluster