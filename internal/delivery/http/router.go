package http

import (
	"toki/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

func NewRouter(itemHandler *handler.ItemHandler, stockHandler *handler.StockHandler, inboundHandler *handler.InboundHandler) *fiber.App {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "OK",
		})
	})

	api := app.Group("/api/v1")
	api.Post("/items", itemHandler.Create)
	api.Get("/items", itemHandler.FindAll)
	api.Get("/items/:id", itemHandler.FindByID)
	api.Put("/items/:id", itemHandler.Update)
	api.Delete("items/:id", itemHandler.Delete)

	api.Get("/stocks/:item_id", stockHandler.GetStock)

	api.Post("/inbounds", inboundHandler.Create)

	return app
}
