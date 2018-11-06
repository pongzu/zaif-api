package main

import (
	"context"
	"fmt"
	"strings"
)

func (c *client) GetPrice(ctx context.Context, pair string) (Data, error) {
	path := fmt.Sprintf("last_price/%s", pair)

	resp, err := c.getResponse(path, ctx)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *client) GetTicker(ctx context.Context, pair string) (Data, error) {
	path := fmt.Sprintf("ticker/%s", pair)

	resp, err := c.getResponse(path, ctx)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (c *client) getResponse(path string, ctx context.Context) (Data, error) {
	rawMsg, err := c.do(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	switch s := strings.Split(path, "/"); s[0] {

	case "last_price":
		var data = new(Price)
		if err := data.encode(rawMsg); err != nil {
			return nil, err
		}
		return data, nil
	case "ticker":
		var data = new(Ticker)
		if err := data.encode(rawMsg); err != nil {
			return nil, err
		}
		return data, nil
	}
	return nil, nil
}
