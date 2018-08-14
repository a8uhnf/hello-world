#!/usr/bin/env bash

docker build -t a8uhnf/kubectl:v1.9.0 ./kubectl
docker push a8uhnf/kubectl:v1.9.0