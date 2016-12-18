package dcap

import (
	"os"
)

func NewLocalDcap(name string, flag int, perm os.FileMode) (*Dcap, error) {

	var d Dcap
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return &d, err
	}

	d = Dcap{Reader: f, Writer: f, Closer: f}
	return &d, nil
}
