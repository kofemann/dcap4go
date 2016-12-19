package dcap

import (
	"os"
)

func NewLocalDcap(name string, flag int, perm os.FileMode) (*DcapStream, error) {

	var d DcapStream
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return &d, err
	}

	d = DcapStream{Reader: f, Writer: f, Closer: f}
	return &d, nil
}
