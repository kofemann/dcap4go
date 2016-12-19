package dcap

import (
	"bufio"
	"io"
	"net"
)

type DcapStream struct {
	Reader     io.Reader
	Writer     io.Writer
	Closer     io.Closer
	ReaderFrom io.ReaderFrom
	WriterTo   io.WriterTo

	dcap *Dcap
}

type Dcap struct {
	conn           net.Conn
	controlChannel *bufio.ReadWriter
	dataChannel    net.Conn
	session        uint32
}
