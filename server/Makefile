.PHONY: proto install
proto:
	protoc -I proto proto/cite/cite.proto --go_out=./proto/ --go_opt=paths=source_relative --go-grpc_out=./proto/ --go-grpc_opt=paths=source_relative

install:
	brew install protobuf
	go get -u github.com/golang/protobuf/protoc-gen-go