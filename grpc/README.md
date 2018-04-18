$ protoc -I api/ \
    -I${GOPATH}/src \
    --go_out=plugins=grpc:api \
    api/api.proto
$ go build -i -v -o bin/server gitlab.com/pantomath-io/demo-grpc/server

$ go build -i -v -o bin/client gitlab.com/pantomath-io/demo-grpc/client