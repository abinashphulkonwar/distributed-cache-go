package api

import (
	"errors"

	"github.com/abinashphulkonwar/dist-cache/api/routes"
	"github.com/abinashphulkonwar/dist-cache/storage"
	"github.com/gofiber/fiber/v2"
)

type errorRes struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func App(db *storage.BadgerStorage) *fiber.App {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error

			if errors.As(err, &e) {
				code = e.Code
			}

			res := errorRes{
				Status: code,
			}

			if err != nil {

				message := err.Error()
				println(message)

				res.Message = message
				return ctx.Status(code).JSON(res)
			}
			res.Message = e.Error()
			return ctx.Status(code).JSON(res)
		},
	})

	routes.WriteRoute(app, db)
	routes.QueryRoute(app, db)

	return app
}
