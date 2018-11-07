package main

import (
	"encoding/json"
)

type Trades []struct {
	Date         int     `json:"date"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	Tid          int     `json:"tid"`
	CurrencyPair string  `json:"currency_pair"`
	TradeType    string  `json:"trade_type"`
}

func (trades *Trades) unmshl(rawMsg json.RawMessage) error {
	if err := json.Unmarshal(rawMsg, trades); err != nil {
		return err
	}
	return nil
}
