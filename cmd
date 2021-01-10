#!/bin/bash

set -e

GOOGLEAPIS_DIR=~/Downloads/googleapis/

# Auto-generate golang bindings for protobuf
mkdir -p calculatorpb
protoc calculator.proto --proto_path=${GOOGLEAPIS_DIR} --go_opt=paths=source_relative --proto_path=. --go_grpc_opt=require_unimplemented_servers=false,paths=source_relative --go_grpc_out=./calculatorpb --go_out=./calculatorpb
