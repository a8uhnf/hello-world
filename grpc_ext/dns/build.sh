#!/usr/bin/env bash

rm grpc_ext_client_dns
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o grpc_ext_client_dns main.go
docker build -t a8uhnf/grpc_ext_client_dns:1.0 .
rm grpc_ext_client_dns
docker push a8uhnf/grpc_ext_client_dns:1.0