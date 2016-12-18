package dcap

import (
	"bufio"
	"fmt"
	"net"
	"net/url"
	"os"
)

func NewRemoteDcap(u *url.URL, flag int, perm os.FileMode) (*Dcap, error) {

	var d Dcap

	conn, err := net.Dial("tcp", u.Host)
	if err != nil {
		return &d, err
	}

	d.conn = conn
	d.controlChannel = bufio.NewReadWriter(bufio.NewReader(d.conn), bufio.NewWriter(d.conn))
	d.asciiHandshake()

	return &d, nil
}

func (d *Dcap) sendAsciiCommand(command string) (string, error) {

	msg := fmt.Sprintf("%d %d client %s\n", d.session, 0, command)
	_, err := d.controlChannel.WriteString(msg)
	if err != nil {
		return "", err
	}

	err = d.controlChannel.Flush()
	if err != nil {
		return "", err
	}

	rep, _, err := d.controlChannel.ReadLine()
	d.session++

	return string(rep), err
}

func (d *Dcap) asciiHandshake() error {

	rep, err := d.sendAsciiCommand("hello 0 0 0 0")
	if err == nil {
		fmt.Println(rep)
	}
	return nil
}
