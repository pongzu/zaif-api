package zaif

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Price struct {
	Price float64 `json:"last_price"`
}

func (price *Price) WriteTo(w io.Writer) {
	bs := float64ToBytes(price.Price)
	w.Write(bs)
}

func float64ToBytes(f float64) []byte {
	var buf bytes.Buffer
	err := binary.Write(&buf, binary.BigEndian, f)
	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	return buf.Bytes()
}
