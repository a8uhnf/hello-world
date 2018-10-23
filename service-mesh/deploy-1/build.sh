#!/usr/bin/env bash

rm deploy-1
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o deploy-1 main.go
docker build -t a8uhnf/envoy-mesh:deploy-1 .
rm deploy-1
docker push a8uhnf/envoy-mesh:deploy-1