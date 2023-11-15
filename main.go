package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/hosseinmirzapur/microservice/client"
	"github.com/hosseinmirzapur/microservice/proto"
)

func main() {
	// cli flags
	jsonMode := flag.Bool("json", true, "true: json | false: grpc")
	clientMode := flag.Bool("client", true, "true: client | false: server")
	listenAddr := flag.String("listen", "0.0.0.0:3000", "listening address")

	flag.Parse()

	if *jsonMode {
		handleJSON(*clientMode, *listenAddr)
	} else {
		handleGRPC(*clientMode, *listenAddr)
	}
}

func handleJSON(clientMode bool, listenAddr string) {
	if clientMode {
		c := client.New(listenAddr)
		res, err := c.FetchPrice(context.Background(), "ETH")
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Printf("price is: %+v\n", res.Price)
	} else {
		svc := NewLoggingService(
			NewMetricService(&priceFetcher{}),
		)

		server := NewJSONAPIServer(listenAddr, svc)

		server.Run()
	}
}

func handleGRPC(clientMode bool, listenAddr string) {
	if clientMode {
		c, err := client.NewGRPC(listenAddr)
		if err != nil {
			log.Fatal(err.Error())
		}

		res, err := c.FetchPrice(context.Background(), &proto.PriceRequest{
			Ticker: "ETH",
		})
		if err != nil {
			log.Fatal(err.Error())
		}

		fmt.Printf("price is: %+v\n", res.Price)

	} else {
		svc := NewLoggingService(
			NewMetricService(&priceFetcher{}),
		)
		err := makeGRPCServerAndRun(listenAddr, svc)
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
