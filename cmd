#!/bin/bash

set -e

GOOGLEAPIS_DIR=~/Downloads/googleapis/

# Auto-generate golang bindings for protobuf
mkdir -p calculatorpb
protoc calculator.proto --proto_path=${GOOGLEAPIS_DIR} --proto_path=. --go_grpc_out=./calculatorpb --go_out=./calculatorpb
