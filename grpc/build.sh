#!/usr/bin/env bash

rm grpc-server
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o grpc-server ./server/main.go
docker build -t a8uhnf/grpc-server:1.0 .
rm grpc-server
docker push a8uhnf/grpc-server:1.0