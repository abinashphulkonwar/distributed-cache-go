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
	"github.com/abinashphulkonwar/dist-cache/tcp/node"
)

func handleRequest(c *node.Connection) {
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
			handlers.ErrorHandler(&handlers.ErrorRes{
				Message: err.Error(),
				Status:  handlers.FAIL,
			}, c)
			continue
		}
		res := handlers.Response{}

		if body.Type == "write" {

			data, err := handlers.WriteDoc(&body, c.DB)
			if err != nil {
				fmt.Println(err)
				handlers.ErrorHandler(&handlers.ErrorRes{
					Message: err.Error(),
					Status:  handlers.FAIL,
				}, c)
				continue
			}

			res.Message = "Data written successfully"

			switch body.Data.Commend {
			case handlers.SET:
				res.Data = [2]string{data.Key, data.Value}
			case handlers.GET:
				res.Data = [2]string{data.Key, "1"}
			case handlers.DELETE:
				res.Data = [2]string{body.Data.Key, "0"}
			}
			res.Status = handlers.SUCCESS

			jsonData, err := json.Marshal(res)
			if err != nil {
				fmt.Println(err)
				handlers.ErrorHandler(&handlers.ErrorRes{
					Message: err.Error(),
					Status:  handlers.FAIL,
				}, c)
				continue
			}
			c.Conn.Write(jsonData)
			continue
		}
		handlers.ErrorHandler(&handlers.ErrorRes{
			Message: "Not found",
			Status:  handlers.FAIL,
		}, c)
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

		client := &node.Connection{
			Conn: conn,
			ID:   id,
			DB:   db,
		}

		println(client.ID)

		node.SetConnectionToMap(client)

		conn.Write([]byte(client.ID))

		go handleRequest(client)
		conn.Write([]byte("Message received.\n"))

	}

}
