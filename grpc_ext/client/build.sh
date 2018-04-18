#!/usr/bin/env bash

rm grpc_ext_client
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o grpc_ext_client main.go
docker build -t a8uhnf/grpc_ext_client:1.0 .
rm grpc_ext_client
docker push a8uhnf/grpc_ext_client:1.0