package tcp

import (
	"net"

	"github.com/abinashphulkonwar/dist-cache/storage"
)

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Status  int    `json:"status"`
}

type Client struct {
	Conn net.Conn
	ID   string
	DB   *storage.BadgerStorage
}

type ErrorRes struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
