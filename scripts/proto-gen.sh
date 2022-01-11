#!/bin/bash

cd protos
protoc -I=. --go-grpc_out=require_unimplemented_servers=false:../pkg --go_out=:../pkg poster.proto
