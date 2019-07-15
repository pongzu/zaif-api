package zaif

import (
	"fmt"
	"io"
)

type Trades struct {
	trades []Trade
}

type Trade struct {
	Date   int     `json:"date"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
	Tid    int     `json:"tid"　`
}

func (trades *Trades) WriteTo(w io.Writer) {
	var bufString string
	for i, trade := range trades.trades {
		if i == 0 {
			bufString += fmt.Sprintf("Date: %v Price: %v Amount: %v Tid: %v", trade.Date, trade.Price, trade.Amount, trade.Tid)
			continue
		}
		bufString += fmt.Sprintf("\nDate: %v Price: %v Amount: %v Tid: %v", trade.Date, trade.Price, trade.Amount, trade.Tid)
	}
	w.Write([]byte(bufString))
}
