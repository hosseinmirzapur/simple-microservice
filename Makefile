build:
	go build -o bin/pricefetcher

server: build
	./bin/pricefetcher

client: build
	./bin/pricefetcher -client

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. \
	--go-grpc_opt=paths=source_relative routeguide/route_guide.proto