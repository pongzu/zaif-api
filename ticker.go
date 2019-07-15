package zaif

import (
	"fmt"
	"io"
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

func (tiker *Ticker) WriteTo(w io.Writer) {
	r := fmt.Sprintf("Last: %v High: %v Low: %v Volume: %v Bid: %v Ask: %", tiker.Last, tiker.High, tiker.Low, tiker.Volume, tiker.Bid, tiker.Ask)
	w.Write([]byte(r))
}
