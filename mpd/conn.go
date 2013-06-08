package mpd

import (
	"bufio"
	"fmt"
	"net"
)

const okStr = "OK\n"

type Conn interface {
	Connected() bool
}

type mpdConn struct {
	conn net.Conn
	connVer string
	events chan Event
	actions chan Action
}

func (c mpdConn) Connected() bool {
	return c.conn != nil
}

func Open(addr string) (Conn, error) {
	// Dial into the port
	conn, err := net.Dial("tcp", addr)

	// Make sure it worked.
	if err != nil {
		return nil, err
	}

	// Read the version line.
	reader := bufio.NewReader(conn)
	connVer, err := reader.ReadString('\n')

	// We couldn't even read that, give up.
	if err != nil {
		return nil, fmt.Errorf("Unable to read from stream.")
	}

	// Ping the server.
	conn.Write([]byte("ping\n"))
	resp, err := reader.ReadString('\n')

	// Make sure we got an OK.
	if err != nil || resp != okStr {
		return nil, fmt.Errorf("Unable to communicate with server.")
	}

	// Give them the mpdConn
	return mpdConn{
		conn,
		connVer,
		make(chan Event),
		make(chan Action),
	}, nil
}
