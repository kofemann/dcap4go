package dcap

import (
	"bufio"
	"io"
	"net"
)

type Dcap struct {
	Reader io.Reader
	Writer io.Writer
	Closer io.Closer

	conn           net.Conn
	controlChannel *bufio.ReadWriter
	dataChannel    net.Conn
	session        uint32
}
