package api_test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/abinashphulkonwar/dist-cache/api"
	"github.com/abinashphulkonwar/dist-cache/api/handlers"
	"github.com/abinashphulkonwar/dist-cache/storage"
	"github.com/dgraph-io/badger/v3"
	"github.com/gofiber/fiber/v2"
)

func Init() *fiber.App {
	connection, err := badger.Open(badger.DefaultOptions("").WithInMemory(true).WithLogger(nil))

	if err != nil {
		panic(err)
	}
	db := storage.NewBadgerStorage(connection)
	app := api.App(db)
	return app
}

func TestApp(t *testing.T) {
	app := Init()

	data, err := json.Marshal(handlers.Body{
		Key:  "key",
		Data: "value",
	})

	if err != nil {
		t.Errorf("Error adding key value pair")
		return
	}

	req := httptest.NewRequest("POST", "/write", bytes.NewReader(data))
	resp, err := app.Test(req)
	if err != nil {
		t.Errorf("Error adding key value pair " + err.Error())
		return
	}

	println(resp.StatusCode)
	body, _ := io.ReadAll(resp.Body)

	println(string(body))
	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Error adding key value pair")
		return
	}

}
