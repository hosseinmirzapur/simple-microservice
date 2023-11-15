build:
	go build -o bin/pricefetcher

json_server: build
	./bin/pricefetcher -json

json_client: build
	./bin/pricefetcher -json -client

grpc_server: build
	./bin/pricefetcher

grpc_client: build
	./bin/pricefetcher -client

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative proto/service.proto

.PHONY: proto