package node

import (
	"errors"
	"net"

	"github.com/abinashphulkonwar/dist-cache/storage"
)

type Client struct {
	Conn net.Conn
	ID   string
	DB   *storage.BadgerStorage
}

var ConnectionMap = map[string]*Client{}

func SetConnectionToMap(c *Client) {
	ConnectionMap[c.ID] = c
}

func GetConnectionFromMap(id string) (*Client, error) {
	data, ok := ConnectionMap[id]

	if !ok {
		return nil, errors.New("connection not found")
	}
	return data, nil
}
