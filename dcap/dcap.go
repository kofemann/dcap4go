package dcap

import (
	"io"
	"net/url"
	"os"
	"syscall"
)

const (
	DCAP_PROTO        = "dcap"
	DCAP_DEFAULT_PORT = "22125"
)

func Open(fname string, flag int, perm os.FileMode) (*DcapStream, error) {

	var d DcapStream

	u, err := url.Parse(fname)
	if err != nil {
		return &d, err
	}

	if len(u.Scheme) == 0 {
		// local file

		return NewLocalDcap(fname, flag, perm)
	}

	if len(u.Scheme) > 0 && u.Scheme != DCAP_PROTO {
		return &d, os.NewSyscallError("Unsupportd protocol ["+u.Scheme+"]", syscall.EINVAL)
	}

	return NewRemoteDcap(u, flag, perm)
}

// implement interface io.Reader
func (d *Dcap) Read(p []byte) (int, error) {
	return 0, nil
}

// implement interface io.Writer
func (d *Dcap) Write(p []byte) (int, error) {
	return 0, nil
}

// implement interface io.ReadFrom
func (d *Dcap) ReadFrom(r io.Reader) (n int64, err error) {
	return 0, nil
}

// implement interface io.WriteTo
func (d *Dcap) WriteTo(w io.Writer) (n int64, err error) {
	return 0, nil
}

// implement interface io.Closer
func (d *Dcap) Close() error {
	return nil
}
