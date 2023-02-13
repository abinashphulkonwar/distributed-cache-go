package storage

import "github.com/dgraph-io/badger/v3"

type Storage interface {
	Add()
}

func Add() {
	badger.Open(badger.DefaultOptions("path/to/db")
}
