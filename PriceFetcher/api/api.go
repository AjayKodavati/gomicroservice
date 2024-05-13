package api

import (
	"PriceFetcher/service"
	"PriceFetcher/types"
	"context"
	"encoding/json"
	"math/rand"
	"net/http"
)

type JSONAPIServer struct {
	listenAddr string
	svc        *service.PriceFetcher
}

type APIFunc func(context.Context, http.ResponseWriter, *http.Request) error

func NewJSONAPIServer(listenAddr string, svc *service.PriceFetcher) *JSONAPIServer {
	return &JSONAPIServer{
		listenAddr: listenAddr,
		svc:        svc,
	}
}

func (s *JSONAPIServer) Run() {
	http.HandleFunc("/", makeHTTPHandlerFunc(s.handleFetchPrice))
	http.ListenAndServe(s.listenAddr, nil)
}

func makeHTTPHandlerFunc(apiFn APIFunc) http.HandlerFunc {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "requestID", rand.Intn(10000000))

	return func(w http.ResponseWriter, r *http.Request) {
		if err := apiFn(ctx, w, r); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"error": err.Error()})
		}
	}
}

func (s *JSONAPIServer) handleFetchPrice(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	ticker := r.URL.Query().Get("ticker")

	price, err := s.svc.FetchPrice(ctx, ticker)

	if err != nil {
		return err
	}

	priceResp := types.PriceResponse{
		Price:  price,
		Ticker: ticker,
	}
	return writeJSON(w, http.StatusOK, &priceResp)

}

func writeJSON(w http.ResponseWriter, statusCode int, data any) error {
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(data)
}
