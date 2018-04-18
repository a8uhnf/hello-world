#!/usr/bin/env bash

rm grpc_ext_server
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o grpc_ext_server main.go
docker build -t a8uhnf/grpc_ext_server:1.0 .
rm grpc_ext_server
docker push a8uhnf/grpc_ext_server:1.0