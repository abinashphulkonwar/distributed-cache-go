package main

import (
	"github.com/abinashphulkonwar/dist-cache/api"
	"github.com/abinashphulkonwar/dist-cache/storage"
	"github.com/dgraph-io/badger/v3"
)

func main() {
	connection, err := badger.Open(badger.DefaultOptions("db"))

	if err != nil {
		panic(err)
	}

	defer connection.Close()

	db := storage.NewBadgerStorage(connection)

	api.ApiServer(db)

}
