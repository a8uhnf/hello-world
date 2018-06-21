#!/usr/bin/env bash

rm logging-ext
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o logging-ext main.go
docker build -t a8uhnf/logging-ext:1.0 .
rm logging-ext
docker push a8uhnf/logging-ext:1.0