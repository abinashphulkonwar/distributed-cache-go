package handlers

import (
	"encoding/json"
	"errors"

	"github.com/abinashphulkonwar/dist-cache/storage"
	"github.com/gofiber/fiber/v2"
)

type Body struct {
	Key  string `json:"Key"`
	Data string `json:"Data"`
}

type response struct {
	Message string `json:"message"`
	Data    string `json:"data"`
	Status  int    `json:"status"`
}

func WriteDoc(c *fiber.Ctx, db *storage.BadgerStorage) error {
	bytes := c.Body()
	if (len(bytes)) == 0 {

		return &fiber.Error{
			Code:    fiber.StatusBadRequest,
			Message: "request body is empty",
		}
	}
	data := Body{}
	err := json.Unmarshal(bytes, &data)
	if err != nil {
		return err
	}
	if data.Key == "" || data.Data == "" {
		return errors.New("key or data is empty")
	}

	err = db.Add(data.Key, data.Data)
	if err != nil {
		return err
	}
	res := response{
		Message: "success",
		Status:  200,
	}
	return c.Status(200).JSON(res)

}
