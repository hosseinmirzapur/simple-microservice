build:
	go build -o bin/pricefetcher

server: build
	./bin/pricefetcher

client: build
	./bin/pricefetcher -client