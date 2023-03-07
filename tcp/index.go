package tcp

import (
	"github.com/abinashphulkonwar/dist-cache/storage"
)

func ApiServer(db *storage.BadgerStorage) {
	err := App(db)
	if err != nil {
		panic(err)
	}
}
