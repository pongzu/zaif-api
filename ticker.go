package main

import (
	"encoding/json"
)

type Ticker struct {
	Last   float64 `json:"last"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Vwap   float64 `json:"vwap"`
	Volume float64 `json:"volume"`
	Bid    float64 `json:"bid"`
	Ask    float64 `json:"ask"`
}

// unmshl allows to unmarshal rawmessage to Ticker structure
func (p *Ticker) unmshl(rawMsg json.RawMessage) error {
	if err := json.Unmarshal(rawMsg, p); err != nil {
		return err
	}
	return nil
}
