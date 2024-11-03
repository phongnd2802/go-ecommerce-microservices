//go:build tools
// +build tools

package tools

import (
	_ "google.golang.org/protobuf/cmd/protoc-gen-go" 
	_ "github.com/sqlc-dev/sqlc"
	_ "github.com/google/wire"
)