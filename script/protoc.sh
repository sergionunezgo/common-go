#!/bin/sh

PROTO_DIR=goreuse/test/v2

protoc --proto_path=${PROTO_DIR} --go_out=. --go-grpc_out=. \
    --go_opt=module=github.com/sergionunezgo/go-reuse/v2 \
    --go-grpc_opt=module=github.com/sergionunezgo/go-reuse/v2 \
    test.proto