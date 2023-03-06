package tcp

import (
	"github.com/abinashphulkonwar/dist-cache/storage"
)

func ApiServer(db *storage.BadgerStorage) error {
	err := App(db)
	if err != nil {
		return err
	}
	return nil
}
