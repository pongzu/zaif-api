package zaif

type Trades struct {
	trades []Trade
}

type Trade struct {
	Date   int     `json:"date"`
	Price  float64 `json:"price"`
	Amount float64 `json:"amount"`
	Tid    int     `json:"tid"　`
}
