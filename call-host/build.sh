#!/usr/bin/env bash

rm call-host
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o call-host main.go
docker build --network="host" -t a8uhnf/call-host:1.0 .
rm call-host
docker push a8uhnf/call-host:1.0