//go:build tools
// +build tools

package tools

import (
	_ "google.golang.org/protobuf/cmd/protoc-gen-go" 
	_ "github.com/sqlc-dev/sqlc"
	_ "github.com/google/wire"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
)