package dcap

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const (
	END_OF_DATA   = -1
	DCAP_WRITE    = 1
	DCAP_READ     = 2
	DCAP_SEEK     = 3
	DCAP_CLOSE    = 4
	DCAP_READV    = 13
	DATA          = 8
	DCAP_SEEK_SET = 0
	DCAP_SEEK_CUR = 1
	DCAP_SEEK_END = 2
)

func NewDcapStream(dcap *Dcap) (*DcapStream, error) {
	stream := DcapStream{Reader: &DcapReader{conn: dcap.dataChannel}}
	return &stream, nil
}

func getAck(r io.Reader) error {

	var l int32
	err := binary.Read(r, binary.BigEndian, &l)
	if err != nil {
		return err
	}

	fmt.Print("Ack: ")
	fmt.Println(l)
	b := make([]byte, l)
	_, err = io.ReadFull(r, b)
	fmt.Print("Ack2: ")
	// FIXME: handle errors
	return err
}

type DcapReader struct {
	conn net.Conn
}

// implement interface io.Reader
func (dcapReader *DcapReader) Read(p []byte) (int, error) {

	binary.Write(dcapReader.conn, binary.BigEndian, int(12))
	binary.Write(dcapReader.conn, binary.BigEndian, int(DCAP_READ))
	binary.Write(dcapReader.conn, binary.BigEndian, uint64(len(p)))

	err := getAck(dcapReader.conn)
	if err != nil {
		return 0, err
	}

	// receice the header
	var l uint32
	err = binary.Read(dcapReader.conn, binary.BigEndian, &l)
	if err != nil {
		return 0, err
	}

	err = binary.Read(dcapReader.conn, binary.BigEndian, &l)
	if err != nil {
		return 0, err
	}

	byteRead := 0
	for byteRead < len(p) {
		var count int32
		err = binary.Read(dcapReader.conn, binary.BigEndian, &count)
		if err != nil {
			return 0, err
		}

		if count == END_OF_DATA {
			getAck(dcapReader.conn)
			break
		}
		n, err := io.ReadFull(dcapReader.conn, p[byteRead:byteRead+int(count)])
		if err != nil {
			return 0, err
		}

		byteRead += n
	}

	return byteRead, nil
}

// implement interface io.Writer
func (d *DcapStream) Write(p []byte) (int, error) {
	return 0, nil
}

// implement interface io.ReadFrom
func (d *DcapStream) ReadFrom(r io.Reader) (n int64, err error) {
	return 0, nil
}

// implement interface io.WriteTo
func (d *DcapStream) WriteTo(w io.Writer) (n int64, err error) {
	return 0, nil
}

// implement interface io.Closer
func (d *DcapStream) Close() error {
	return nil
}
