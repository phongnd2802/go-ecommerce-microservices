#!/usr/bin/env bash

# Xác định đường dẫn thư mục hiện tại
DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

# Tạo thư mục pb nếu chưa có
mkdir -p ${DIR}/pb
# Xóa sạch các tệp cũ trong thư mục pb
rm -rf ${DIR}/pb/*

# Tìm và biên dịch tất cả các tệp .proto trong thư mục proto và các thư mục con
proto_files=$(find "${DIR}/proto" -name '*.proto')

protoc -I ${DIR}/proto \
  --go_out=${DIR}/pb --go_opt=paths=source_relative \
  --go-grpc_out=${DIR}/pb --go-grpc_opt=paths=source_relative \
  ${proto_files}
