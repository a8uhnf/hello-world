#!/usr/bin/env bash

rm go-hit
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go-hit main.go
docker build -t a8uhnf/go-hit:1.0 .
rm go-hit
docker push a8uhnf/go-hit:1.0