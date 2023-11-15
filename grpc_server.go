package main

import (
	"context"
	"net"

	"github.com/hosseinmirzapur/microservice/proto"
	"google.golang.org/grpc"
)

func makeGRPCServerAndRun(listenAddr string, svc PriceFetcher) error {
	grpcPriceFetcher := NewGRPCPriceFetcherServer(svc)

	ln, err := net.Listen("tls", listenAddr)
	if err != nil {
		return err
	}

	opts := []grpc.ServerOption{}
	server := grpc.NewServer(opts...)

	proto.RegisterPriceFetcherServer(server, grpcPriceFetcher)

	return server.Serve(ln)
}

type GRPCPriceFetcherServer struct {
	svc PriceFetcher
	proto.UnimplementedPriceFetcherServer
}

func NewGRPCPriceFetcherServer(svc PriceFetcher) *GRPCPriceFetcherServer {
	return &GRPCPriceFetcherServer{svc: svc}
}

func (s *GRPCPriceFetcherServer) FetchPrice(ctx context.Context, req *proto.PriceRequest) (*proto.PriceResponse, error) {
	price, err := s.svc.FetchPrice(ctx, req.Ticker)
	if err != nil {
		return nil, err
	}
	resp := &proto.PriceResponse{Ticker: req.Ticker, Price: float32(price)}
	return resp, nil
}
