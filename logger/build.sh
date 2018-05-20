#!/usr/bin/env bash

rm logger
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o logger main.go
docker build -t a8uhnf/logger:1.0 .
rm logger
docker push a8uhnf/logger:1.0