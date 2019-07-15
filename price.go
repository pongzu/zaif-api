package zaif

import (
	"io"
	"strconv"
)

type Price struct {
	Price float64 `json:"last_price"`
}

func (price *Price) WriteTo(w io.Writer) {
	i := int(price.Price)
	bs := []byte(strconv.Itoa(i) + "円")
	w.Write(bs)
}
