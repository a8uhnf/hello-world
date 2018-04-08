#!/usr/bin/env bash

rm server
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o server main.go
docker build -t a8uhnf/grpc-server:1.0 .
rm server
docker push a8uhnf/grpc-server:1.0