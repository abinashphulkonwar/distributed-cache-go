package handlers

import (
	"github.com/abinashphulkonwar/dist-cache/storage"
	"github.com/gofiber/fiber/v2"
)

func QueryDoc(c *fiber.Ctx, db *storage.BadgerStorage) error {
	key := c.Query("key")
	if key == "" {
		return &fiber.Error{
			Code:    fiber.StatusUnprocessableEntity,
			Message: "query key is empty",
		}
	}
	data, err := db.Get(key)

	if err != nil {
		return err
	}

	if data[0] == "" || data[1] == "" {
		return &fiber.Error{
			Code:    fiber.StatusNotFound,
			Message: "not found!",
		}
	}

	res := Response{
		Message: "success",
		Status:  200,
		Data:    data,
	}
	return c.Status(200).JSON(res)

}
