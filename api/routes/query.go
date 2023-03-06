package routes

import (
	"github.com/abinashphulkonwar/dist-cache/api/handlers"
	"github.com/abinashphulkonwar/dist-cache/storage"
	"github.com/gofiber/fiber/v2"
)

func QueryRoute(c *fiber.App, db *storage.BadgerStorage) {

	c.Get("/query", func(c *fiber.Ctx) error {
		return handlers.QueryDoc(c, db)
	})

}
