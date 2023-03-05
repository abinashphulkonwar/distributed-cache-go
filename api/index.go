package api

import (
	"github.com/abinashphulkonwar/dist-cache/storage"
)

func ApiServer(db *storage.BadgerStorage) {
	app := App(db)
	app.Listen(":3000")
}
