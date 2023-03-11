package node

import (
	"net"
)

func (c *Connection) SetConnection() {
	ConnectionMap[c.ID] = c
}

func ConnectionHandler(ip string) (*Connection, error) {
	conn, err := net.Dial("tcp", ip)
	if err != nil {
		return nil, err
	}

	return &Connection{
		ID:   ip,
		Conn: conn,
	}, nil

}
