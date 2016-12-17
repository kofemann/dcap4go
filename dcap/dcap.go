package dcap

import (
	"errors"
	"fmt"
	"net/url"
)

func Open(fname string) (Dcap, error) {

	var d Dcap

	u, err := url.Parse(fname)
	if err != nil {
		return d, err
	}

	fmt.Println(u)
	return d, errors.New("not implemented")
}

func (d *Dcap) Close() {

}
