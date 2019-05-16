#!/usr/bin/env bash

if [ -z $1 ]
then
    echo "Need to provide image tag..."
    exit
fi

rm go
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o go main.go
docker build -t a8uhnf/go:$1 .
rm go
docker push a8uhnf/go:$1