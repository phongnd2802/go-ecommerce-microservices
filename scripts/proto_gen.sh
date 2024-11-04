#!/usr/bin/env bash


DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )/.." >/dev/null 2>&1 && pwd )"

rm -rf ${DIR}/pb/*
rm -rf ${DIR}/docs/swagger/*.swagger.json

proto_files=$(find "${DIR}/proto" -name '*.proto' -not -path "*/google/*" -not -path "*/protoc-gen-openapiv2/*")

protoc -I ${DIR}/proto \
  --go_out=${DIR}/pb --go_opt=paths=source_relative \
  --go-grpc_out=${DIR}/pb --go-grpc_opt=paths=source_relative \
  --grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
  --openapiv2_out=docs/swagger --openapiv2_opt=allow_merge=true,merge_file_name=ecommerce_api \
  ${proto_files}

statik -f -src=./docs/swagger -dest=./docs