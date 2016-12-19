package dcap

import (
	"bufio"
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"net/url"
	"os"
	"strings"
)

func NewRemoteDcap(u *url.URL, flag int, perm os.FileMode) (*DcapStream, error) {

	var d DcapStream

	conn, err := net.Dial("tcp", u.Host)
	if err != nil {
		return &d, err
	}

	client := Dcap{}
	client.conn = conn
	client.controlChannel = bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn))
	err = client.asciiHandshake()
	if err != nil {
		return &d, err
	}

	var mode = "r"
	if flag&(os.O_WRONLY|os.O_RDWR) != 0 {
		mode = "w"
	}

	err = client.open(u.String(), mode)
	if err != nil {
		return &d, err
	}
	return NewDcapStream(&client)
}

func parseReply(msg string) ([]string, error) {
	s := strings.Split(msg, " ")
	if s[3] == "failed" {
		return s[3:], errors.New(strings.Join(s[5:], " "))
	}

	return s[3:], nil
}

func (d *Dcap) sendAsciiCommand(command string) (uint32, []string, error) {

	session := d.session
	msg := fmt.Sprintf("%d %d client %s\n", session, 0, command)
	_, err := d.controlChannel.WriteString(msg)
	if err != nil {
		return 0, make([]string, 0), err
	}

	err = d.controlChannel.Flush()
	if err != nil {
		return 0, make([]string, 0), err
	}

	rep, _, err := d.controlChannel.ReadLine()
	d.session++

	if err != nil {
		return 0, make([]string, 0), err
	}

	args, err := parseReply(string(rep))
	return session, args, err
}

func (d *Dcap) asciiHandshake() error {

	_, _, err := d.sendAsciiCommand("hello 0 0 0 0")
	if err != nil {
		return err
	}

	return nil
}

func (d *Dcap) open(u string, mode string) error {

	openMsg := fmt.Sprintf("open %s %s dead.end 0 -passive -uid=%d -gid=%d", u, mode, os.Geteuid(), os.Getgid())
	session, rep, err := d.sendAsciiCommand(openMsg)
	if err != nil {
		return err
	}

	host := rep[1]
	port := rep[2]
	rawChallange := []byte(rep[3])

	fmt.Println(net.JoinHostPort(host, port))
	conn, err := net.Dial("tcp", net.JoinHostPort(host, port))
	if err != nil {
		return err
	}

	d.dataChannel = conn

	binary.Write(conn, binary.BigEndian, int32(session))
	binary.Write(conn, binary.BigEndian, int32(len(rawChallange)))
	conn.Write(rawChallange)

	return nil
}
