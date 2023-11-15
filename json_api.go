package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/hosseinmirzapur/microservice/types"
)

type JSONAPIServer struct {
	listenAddr string
	svc        PriceFetcher
}

func NewJSONAPIServer(listenAddr string, svc PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/price", makeHttpHandlerFunc(s.handleFetchPrice))

	log.Printf("TCP server up and running on: %s", s.listenAddr)
	log.Fatal(http.ListenAndServe(s.listenAddr, nil))
}

func makeHttpHandlerFunc(apiFunc APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", uuid.New().String())
	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFunc(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"err": err.Error()})
		}
	}
}

func (s *JSONAPIServer) handleFetchPrice(
	ctx context.Context,
	w http.ResponseWriter,
	r *http.Request,
) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)
	if err != nil {
		return err
	}

	priceRes := types.PriceResponse{
		Price:  price,
		Ticker: ticker,
	}

	return writeJSON(w, http.StatusOK, &priceRes)
}

func writeJSON(w http.ResponseWriter, s int, v any) error {
	w.WriteHeader(s)

	return json.NewEncoder(w).Encode(v)
}
