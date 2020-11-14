package openflow

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"net"
)

type conn struct {
	MessageIn  chan Message
	MessageOut chan Message

	conn *net.TCPConn
}

func (c *conn) Run() error {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:6333")
	if err != nil {
		return fmt.Errorf("resolve addr: %v", err)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return fmt.Errorf("dial: %v", err)
	}

	r := bufio.NewReader(conn)

	header := make([]byte, 8)

	// TODO: check if connection has been closed
	for {
		_, err := r.Read(header)
		if err != nil {
			return fmt.Errorf("read header: %v", err)
		}

		version := header[0]
		packetType := header[1]
		packetLength := binary.BigEndian.Uint16(header[2:4])
		xid := binary.BigEndian.Uint32(header[5:])
		payload := make([]byte, packetLength-header)

		_, err := r.Read(payload)
		if err != nil {
			return fmt.Errorf("read payload: %v", err)
		}

		c.MessageIn <- Message{
			Header{
				Version: version,
				Type:    packetType,
				Length:  packetLength,
				XID:     xid,
			},
			Payload: payload,
		}
	}

	return nil
}
