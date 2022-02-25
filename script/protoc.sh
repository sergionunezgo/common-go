#!/bin/sh

PROTO_DIR=goreuse/test

protoc --proto_path=${PROTO_DIR} --go_out=. --go-grpc_out=. \
    --go_opt=module=github.com/sergionunezgo/go-reuse \
    --go-grpc_opt=module=github.com/sergionunezgo/go-reuse \
    test.proto