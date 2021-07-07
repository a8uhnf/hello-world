#!/bin/bash

echo "starting streaming server"

export GRPC_ADDRESS=":9000"

GO111MODULE=on go run grpc-streaming/client/client.go