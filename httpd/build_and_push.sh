#!/bin/bash



docker build . --platform "linux/amd64" -t docker.io/a8uhnf/httpd:v1.0.0

docker push docker.io/a8uhnf/httpd:v1.0.0