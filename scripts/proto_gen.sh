#!/usr/bin/env bash


DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." >/dev/null 2>&1 && pwd )"

rm -rf ${DIR}/pb/*

proto_files=$(find "${DIR}/proto" -name '*.proto' -not -path "*/google/*")

protoc -I ${DIR}/proto \
  --go_out=${DIR}/pb --go_opt=paths=source_relative \
  --go-grpc_out=${DIR}/pb --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
  ${proto_files}
