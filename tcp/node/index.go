package node

import (
	"errors"
	"net"

	"github.com/abinashphulkonwar/dist-cache/storage"
)

type Connection struct {
	Conn net.Conn
	ID   string
	DB   *storage.BadgerStorage
}

var ConnectionMap = map[string]*Connection{}

func SetConnectionToMap(c *Connection) {
	ConnectionMap[c.ID] = c
}

func GetConnectionFromMap(id string) (*Connection, error) {
	data, ok := ConnectionMap[id]

	if !ok {
		return nil, errors.New("connection not found")
	}
	return data, nil
}
