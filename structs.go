package zaif-api


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

type Price struct {
	Price float64 `json:"last_price"`
}

type Ticker struct {
	Last   float64 `json:"last"`
	High   float64 `json:"high"`
	Low    float64 `json:"low"`
	Vwap   float64 `json:"vwap"`
	Volume float64 `json:"volume"`
	Bid    float64 `json:"bid"`
	Ask    float64 `json:"ask"`
}

type Trades []struct {
	Date         int     `json:"date"`
	Price        float64 `json:"price"`
	Amount       float64 `json:"amount"`
	Tid          int     `json:"tid"`
	CurrencyPair string  `json:"currency_pair"`
	TradeType    string  `json:"trade_type"`
}
