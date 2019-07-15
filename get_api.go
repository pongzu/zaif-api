package zaif

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

func (c *client) GetPairs(ctx context.Context) (Res, error) {
	path := fmt.Sprintf("currency_pairs/all")

	jsonData, err := c.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "client.do failed")
	}

	pairs := new(Pairs)
	if err := json.Unmarshal(jsonData, &pairs.pairs); err != nil {
		return nil, errors.Wrap(err, "failed to json.Unmarshal")
	}
	return pairs, nil
}

func (c *client) GetPrice(ctx context.Context, pair string) (Res, error) {
	path := fmt.Sprintf("last_price/%s", pair)

	jsonData, err := c.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "client.do failed")
	}
	price := new(Price)

	if err := json.Unmarshal(jsonData, price); err != nil {
		return nil, errors.Wrap(err, "failed to json.Unmarshal")
	}

	return price, nil
}

func (c *client) GetTicker(ctx context.Context, pair string) (Res, error) {
	path := fmt.Sprintf("ticker/%s", pair)

	jsonData, err := c.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "client.do failed")
	}

	tiker := new(Ticker)
	if err := json.Unmarshal(jsonData, tiker); err != nil {
		return nil, errors.Wrap(err, "failed to json.Unmarshal")
	}

	return tiker, nil
}

func (c *client) GetTrades(ctx context.Context, pair string) (Res, error) {
	path := fmt.Sprintf("trades/%s", pair)

	jsonData, err := c.do(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, errors.Wrap(err, "client.do failed")
	}

	trades := new(Trades)
	if err := json.Unmarshal(jsonData, &trades.trades); err != nil {
		return nil, errors.Wrap(err, "failed to json.Unmarshal")
	}
	return trades, nil
}
