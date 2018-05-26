#!/usr/bin/env bash

rm go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go main.go
docker build -t a8uhnf/go:2.0 .
rm go
docker push a8uhnf/go:2.0