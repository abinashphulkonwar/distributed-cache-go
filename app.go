package main

import (
	"github.com/abinashphulkonwar/dist-cache/api"
	"github.com/abinashphulkonwar/dist-cache/storage"
	"github.com/abinashphulkonwar/dist-cache/tcp"
	"github.com/dgraph-io/badger/v3"
)

func main() {

	connection, err := badger.Open(badger.DefaultOptions("db"))

	if err != nil {
		panic(err)
	}

	defer connection.Close()

	db := storage.NewBadgerStorage(connection)

	go api.ApiServer(db)
	println("api server started")
	tcp.ApiServer(db)
	println("tcp server started")

}
