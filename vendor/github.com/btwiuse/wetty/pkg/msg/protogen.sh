#!/bin/bash

protoc \
  --plugin=protoc-gen-ts=$(which protoc-gen-ts) \
  -I . \
  --js_out=import_style=commonjs,binary:. \
  --ts_out=service=grpc-web:. \
  types.proto

mv *.ts *.js ../../js/msg/

cd .. && go generate
