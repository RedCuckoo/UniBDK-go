#!/bin/bash
protoc -I./proto/definitions --go_out=./proto/generated/proto --go-grpc_out=./proto/generated/proto --grpc-gateway_out=./proto/generated/proto --grpc-gateway_opt logtostderr=true proto/definitions/*
