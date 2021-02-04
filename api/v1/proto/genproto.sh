protoc -I . \
      -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
      --go_out . --go_opt paths=source_relative \
      --go-grpc_out . --go-grpc_opt paths=source_relative \
      model_service.proto
protoc -I/usr/local/include -I. \
    -I $GOPATH/src \
    -I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
    --grpc-gateway_out=logtostderr=true:. model_service.proto