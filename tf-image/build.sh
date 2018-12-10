#!/usr/bin/env bash

rm tf-cicd
#CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sso-app main.go
docker build -t a8uhnf/tf-cicd:1.0 .
rm tf-cicd
docker push a8uhnf/tf-cicd:1.0