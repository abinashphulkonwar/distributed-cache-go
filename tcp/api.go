package tcp

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"github.com/abinashphulkonwar/dist-cache/service"
	"github.com/abinashphulkonwar/dist-cache/storage"
	"github.com/abinashphulkonwar/dist-cache/tcp/handlers"
)

func (c *Client) handleRequest() {
	reader := bufio.NewReader(c.Conn)
	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			c.Conn.Close()
			return
		}
		body := handlers.Body{}

		err = json.Unmarshal([]byte(message), &body)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if body.Type == "write" {
			err = handlers.WriteDoc(&body, c.DB)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}

		fmt.Printf("Message incoming: %s", string(message))
		c.Conn.Write([]byte("Message received.\n"))

	}
}

func App(db *storage.BadgerStorage) error {
	tcp, err := net.Listen("tcp", ":3001")
	if err != nil {
		return errors.New("error")
	}
	defer tcp.Close()

	for {
		conn, err := tcp.Accept()
		if err != nil {
			println("Error accepting: ", err.Error())
			conn.Close()
		}

		id, err := service.RandomUUID()
		if err != nil {
			println("Error generating UUID", err.Error())
			conn.Close()
		}

		client := &Client{
			Conn: conn,
			ID:   id,
			DB:   db,
		}

		println(client.ID)

		conn.Write([]byte(client.ID))

		go client.handleRequest()
		conn.Write([]byte("Message received.\n"))

	}

}
