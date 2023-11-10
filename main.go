package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/hosseinmirzapur/microservice/client"
)

func main() {
	listenAddr := flag.String("listen", "0.0.0.0:3000", "listen address")
	isClient := flag.Bool("client", false, "If provided, you will subscribe to grpc client")
	flag.Parse()

	if !*isClient {
		svc := NewLoggingService(
			NewMetricService(&priceFetcher{}),
		)

		server := NewJSONAPIServer(*listenAddr, svc)

		server.Run()
	} else {
		client := client.New("http://localhost:3000")
		response, err := client.FetchPrice(context.Background(), "ETH")
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Printf("price is: %+v\n", response.Price)
	}
}
