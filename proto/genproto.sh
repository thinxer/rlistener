#!/bin/bash

protoc --go_out=plugins=grpc:$GOPATH/src rlistener.proto
