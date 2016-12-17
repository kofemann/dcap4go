package dcap

import (
	"io"
)

type Dcap struct {
	Reader io.Reader
	Writer io.Writer
}
