package tcp

import "net"

type Body struct {
	Key  string `json:"Key"`
	Data string `json:"Data"`
}

type Response struct {
	Message string `json:"message"`
	Data    any    `json:"data"`
	Status  int    `json:"status"`
}

type Client struct {
	conn net.Conn
	id   string
}

type ErrorRes struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}
