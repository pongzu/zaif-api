package zaif

import (
	"fmt"
	"io"
)

type Pairs struct {
	pairs []Pair
}

type Pair struct {
	Name string `json:"name"`
}

func (pairs *Pairs) WriteTo(w io.Writer) {
	var bufString string
	for i, pair := range pairs.pairs {
		if i == 0 {
			bufString += fmt.Sprintf("%d: %s", i, pair.Name)
			continue
		}
		bufString += fmt.Sprintf("\n%d: %s", i, pair.Name)
	}
	w.Write([]byte(bufString))
}
