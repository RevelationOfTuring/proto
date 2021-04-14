#!/bin/sh

cd /Users/oker/go/src/proto/protobuf
go run proto3_add_token.go
go run json_add_token.go
go run proto3_list_token.go
go run json_list_token.go
