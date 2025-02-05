#!/usr/bin/env bash

file_path=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

protoc \
    --proto_path=${file_path}/proto \
    --go_out=${file_path}/proto --go_opt=paths=source_relative \
    --go-grpc_out=${file_path}/proto --go-grpc_opt=paths=source_relative \
    ${file_path}/proto/grpc.proto
