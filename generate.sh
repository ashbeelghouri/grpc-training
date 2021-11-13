#!/bin/bash
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative greet/greetpb/greet.proto


go run greet/greet_server/server.go
go run greet/greet_client/client.go