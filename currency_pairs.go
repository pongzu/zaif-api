package main

import (
	"encoding/json"
)

type Pairs []struct {
	IsToken      bool    `json:"is_token"`
	Name         string  `json:"name"`
	ItemUnitStep float64 `json:"item_unit_step"`
	AuxUnitPoint int     `json:"aux_unit_point"`
	ItemJapanese string  `json:"item_japanese"`
	EventNumber  int     `json:"event_number"`
	AuxJapanese  string  `json:"aux_japanese"`
	CurrencyPair string  `json:"currency_pair"`
	Seq          int     `json:"seq"`
	Description  string  `json:"description"`
	AuxUnitStep  float64 `json:"aux_unit_step"`
	ItemUnitMin  float64 `json:"item_unit_min"`
	AuxUnitMin   float64 `json:"aux_unit_min"`
	Title        string  `json:"title"`
	ID           int     `json:"id"`
}

func (pairs *Pairs) unmshl(rawMsg json.RawMessage) error {
	if err := json.Unmarshal(rawMsg, pairs); err != nil {
		return err
	}
	return nil
}
