package tcp

import (
	"errors"
	"log"
	"net"

	"github.com/abinashphulkonwar/dist-cache/storage"
)

func App(db *storage.BadgerStorage) error {
	tcp, err := net.Listen("tcp", ":3001")
	if err != nil {
		return errors.New("error")
	}
	defer tcp.Close()

	for {
		conn, err := tcp.Accept()
		if err != nil {
			log.Fatal(err)
		}

		client := &Client{
			conn: conn,
			id:   "1",
		}

		println(client.id)

	}

}
