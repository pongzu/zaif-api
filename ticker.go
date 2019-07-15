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

func (ticker *Ticker) WriteTo(w io.Writer) {
	r := fmt.Sprintf("Last: %v円 High: %v Low円: %v円 Volume: %v円 Bid: %v円 Ask: %v円", int(ticker.Last), int(ticker.High), int(ticker.Low), int(ticker.Volume), int(ticker.Bid), int(ticker.Ask))
	w.Write([]byte(r))
}
