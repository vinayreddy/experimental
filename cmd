#!/bin/bash

set -e

if [ "$1" = "clean" ]; then
  rm -rf calculatorpb
  rm client/calculator_grpc_web_pb.js client/calculator_pb.js
fi

GOOGLEAPIS_DIR=/Users/vinayreddy/code/googleapis

# Auto-generate golang bindings for protobuf
# mkdir -p calculatorpb
# protoc calculator.proto --proto_path=${GOOGLEAPIS_DIR} --go_opt=paths=source_relative --proto_path=. --go_grpc_opt=require_unimplemented_servers=false,paths=source_relative --go_grpc_out=./calculatorpb --go_out=./calculatorpb

# Create javascript for grpc-web
# protoc calculator.proto --proto_path=${GOOGLEAPIS_DIR} --proto_path=. --js_out=import_style=commonjs,binary:./client --grpc-web_out=import_style=commonjs,mode=grpcwebtext:./client

# cd client
# npm install
# npx webpack --debug client.js
# cd -

protoc --proto_path=${GOOGLEAPIS_DIR} -I. --include_imports --include_source_info \
--descriptor_set_out=proto.pb proto/calculator.proto
