# Protobuf compilation examples
Compiling protobuf is not trivial. This repo contains a sequence of increasingly complex
test cases, illustrating different aspects you need to be aware of.

## Running the examples
Install `protoc` and clone the repo.

## Setting up the vendor folder
This is not necessary when cloning the repo, but 
```
go mod init github.com/backlin/prototest
go get github.com/gogo/protobuf/proto
go get google.golang.org/genproto
go mod vendor
```