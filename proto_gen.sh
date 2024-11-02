#!/usr/bin/env bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# rm -rf ${DIR}/pb/*

for dir in $(find ${DIR}/proto -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq); do
  files=$(find "${dir}" -name '*.proto')

  protoc -I ${DIR} --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative ${files}
done