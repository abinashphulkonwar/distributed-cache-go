package storage_test

import (
	"testing"

	"github.com/abinashphulkonwar/dist-cache/storage"
	"github.com/dgraph-io/badger/v3"
)

func Init() *storage.BadgerStorage {
	connection, err := badger.Open(badger.DefaultOptions("").WithInMemory(true).WithLogger(nil))

	if err != nil {
		panic(err)
	}
	db := storage.NewBadgerStorage(connection)

	return db
}

var key = "key"
var value = "value"

func TestAdd(t *testing.T) {
	db := Init()

	err := db.Add(key, value)

	if err != nil {
		t.Errorf("Error adding key value pair")
		return
	}

	db.Close()
}

func TestGet(t *testing.T) {
	db := Init()
	err := db.Add(key, value)

	if err != nil {
		t.Errorf("Error adding key value pair")
		return
	}
	data, err := db.Get(key)
	if err != nil {
		println(err.Error())
		t.Errorf("Error getting key value pair")
		return

	}
	println(data[0], data[1])
	if data[0] != key || data[1] != value {
		t.Errorf("Error getting key value pair")
		return
	}
	db.Close()
}
