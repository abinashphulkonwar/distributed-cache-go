package storage

import "github.com/dgraph-io/badger/v3"

type Storage interface {
	Add()
	Get()
	Close()
}

type BadgerStorage struct {
	db *badger.DB
}

func NewBadgerStorage(db *badger.DB) *BadgerStorage {
	return &BadgerStorage{db: db}
}

func (engine *BadgerStorage) Close() {
	engine.db.Close()
}

func (engine *BadgerStorage) Add(key string, value string) error {

	txn := engine.db.NewTransaction(true)
	defer txn.Discard()
	err := txn.Set([]byte(key), []byte(value))
	if err != nil {
		return err
	}
	err = txn.Commit()

	if err != nil {
		return err
	}
	return nil
}

func (engine *BadgerStorage) Get(key string) ([2]string, error) {
	var value string
	err := engine.db.View(func(txn *badger.Txn) error {

		item, err := txn.Get([]byte(key))
		if err != nil {
			return err
		}
		err = item.Value(func(val []byte) error {
			value = string(val)
			return nil
		})
		return err
	})
	if err != nil {
		return [2]string{"", ""}, err
	}
	return [2]string{key, value}, nil
}

func (engine *BadgerStorage) Delete(key string) error {
	txn := engine.db.NewTransaction(true)
	defer txn.Discard()
	err := txn.Delete([]byte(key))
	if err != nil {
		return err
	}
	err = txn.Commit()

	if err != nil {
		return err
	}
	return nil
}
