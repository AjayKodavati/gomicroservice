package client

import (
	"PriceFetcher/types"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Client struct {
	endPoint string
}

func NewgRpcClient(endPoint string) *Client {
	return &Client{
		endPoint: endPoint,
	}
}

func (c *Client) FetchPrice(ctx context.Context, ticker string) (*types.PriceResponse, error) {
	endPoint := fmt.Sprintf("%s/?ticker=%s", c.endPoint, ticker)

	req, err := http.NewRequest("get", endPoint, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		var httpErr map[string]any
		if err := json.NewDecoder(resp.Body).Decode(&httpErr); err != nil {
			return nil, err
		}
		fmt.Println(err, json.NewDecoder(resp.Body).Decode(&httpErr))
		return nil, fmt.Errorf("service responded with non ok status code: %s", err)
	}

	priceResp := new(types.PriceResponse)
	if err := json.NewDecoder(resp.Body).Decode(priceResp); err != nil {
		return nil, err
	}

	return priceResp, nil
}