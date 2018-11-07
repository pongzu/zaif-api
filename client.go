package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// BaseUrl is an endpoint of api
const BaseUrl = "https://api.zaif.jp/api/1/%s"

// Client interfaces is an interface for client
type Client interface {
	GetPairs(ctx context.Context) (Data, error)
	GetPrice(ctx context.Context, pair string) (Data, error)
	GetTicker(ctx context.Context, pair string) (Data, error)
	GetTrades(ctx context.Context, pair string) (Data, error)
}

type client struct{}

// New cliates a new client
func New() Client {
	return &client{}
}

func (c *client) do(ctx context.Context, method, path string, body io.Reader) (json.RawMessage, error) {
	url := fmt.Sprintf(BaseUrl, path)

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	// header and context are set
	req = req.WithContext(ctx)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{Timeout: time.Duration(10) * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// status code is checked
	if resp.StatusCode != http.StatusOK {
		format := "can not send HTTP reqest with statuscode %d: %s"
		msg, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			msg = []byte("no message came from zaif_server")
		}
		return nil, errors.Errorf(format, resp.StatusCode, msg)
	}

	var rawMsg json.RawMessage
	if err := json.NewDecoder(resp.Body).Decode(&rawMsg); err != nil {
		return nil, err
	}
	return rawMsg, nil
}
