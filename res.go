package zaif

import (
	"io"
)

type Res interface {
	WriteTo(io.Writer)
}
