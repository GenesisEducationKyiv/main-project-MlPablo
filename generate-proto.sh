#!/bin/bash

CreateCurrency() {
  echo "-> Processing: currency.proto"
  protoc -I="./api/proto" \
    --go_out="./" \
    --go_opt=Mcurrency.proto="$1/api/proto/grpc_currency_service" \
    --go-grpc_out=require_unimplemented_servers=false:"./" \
    --go-grpc_opt=Mcurrency.proto="$1/api/proto/grpc_currency_service" \
    --experimental_allow_proto3_optional \
    currency.proto
}

echo "Create services for currency microservice"
CreateCurrency "currency"

