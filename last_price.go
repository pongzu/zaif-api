package main

import (
	"encoding/json"
)

type Price struct {
	Price float64 `json:"last_price"`
}

// unmshl allows to unmarshal rawmessage to Price structure
func (p *Price) unmshl(rawMsg json.RawMessage) error {
	if err := json.Unmarshal(rawMsg, p); err != nil {
		return err
	}
	return nil
}
