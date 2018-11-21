#!/usr/bin/env bash

rm sso-app
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o sso-app main.go
docker build -t a8uhnf/sso-app:1.0 .
rm sso-app
docker push a8uhnf/sso-app:1.0