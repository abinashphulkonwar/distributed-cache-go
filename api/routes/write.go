package routes

import (
	"github.com/abinashphulkonwar/dist-cache/api/handlers"
	"github.com/abinashphulkonwar/dist-cache/storage"
	"github.com/gofiber/fiber/v2"
)

func WriteRoute(c *fiber.App, db *storage.BadgerStorage) {

	c.Post("/write", func(c *fiber.Ctx) error {
		return handlers.WriteDoc(c, db)
	})

}
