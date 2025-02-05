#!/usr/bin/env bash
file_path=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )

protoc --go_out=${file_path} --go_opt=paths=source_absolute \
    --go-grpc_out=${file_path} --go-grpc_opt=paths=source_absolute \
    ${file_path}/proto/vendor.proto
