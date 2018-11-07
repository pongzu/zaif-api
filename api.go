package main

import (
	"context"
	"fmt"
	"strings"
)

func (c *client) GetPrice(ctx context.Context, pair string) (Data, error) {
	path := fmt.Sprintf("last_price/%s", pair)

	data, err := c.getResponse(path, ctx)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c *client) GetTicker(ctx context.Context, pair string) (Data, error) {
	path := fmt.Sprintf("ticker/%s", pair)

	data, err := c.getResponse(path, ctx)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (c *client) getResponse(path string, ctx context.Context) (Data, error) {
	rawMsg, err := c.do(ctx, "GET", path, nil)
	if err != nil {
		return nil, err
	}

	var data Data
	switch s := strings.Split(path, "/"); s[0] {
	case "last_price":
		data = new(Price)
		if err := data.unmshl(rawMsg); err != nil {
			return nil, err
		}
	case "ticker":
		data = new(Ticker)
		if err := data.unmshl(rawMsg); err != nil {
			return nil, err
		}
	}
	return data, nil
}
