package client

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hosseinmirzapur/microservice/proto"
	"github.com/hosseinmirzapur/microservice/types"
	"google.golang.org/grpc"
)

type Client struct {
	endpoint string
}

func New(endpoint string) *Client {
	return &Client{endpoint}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endpoint := fmt.Sprintf("%s/price?ticker=%s", c.endpoint, ticker)
	req, err := http.NewRequest("get", endpoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	priceResponse := &types.PriceResponse{}

	if err := json.NewDecoder(resp.Body).Decode(priceResponse); err != nil {
		return nil, err
	}

	return priceResponse, nil
}

func NewGRPC(remoteAddr string) (proto.PriceFetcherClient, error) {
	conn, err := grpc.Dial(remoteAddr)
	if err != nil {
		return nil, err
	}

	return proto.NewPriceFetcherClient(conn), nil
}
