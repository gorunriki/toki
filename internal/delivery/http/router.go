package http

import (
	"toki/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"

	middleware "toki/internal/delivery/http/middleware"
)

func NewRouter(itemHandler *handler.ItemHandler, stockHandler *handler.StockHandler, inboundHandler *handler.InboundHandler, salesHandler *handler.SalesHandler, reportHandler *handler.ReportHandler, authHandler *handler.AuthHandler) *fiber.App {
	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "OK",
		})
	})

	api := app.Group("/api/v1")
	// api.Post("/items", itemHandler.Create)
	api.Get("/items", itemHandler.FindAll)
	api.Get("/items/:id", itemHandler.FindByID)
	api.Put("/items/:id", itemHandler.Update)
	api.Delete("items/:id", itemHandler.Delete)

	api.Get("/stocks/:item_id", stockHandler.GetStock)

	// api.Post("/inbounds", inboundHandler.Create)

	// api.Post("/sales", salesHandler.Create)

	api.Get("/reports/stocks", reportHandler.Stocks)
	api.Get("/reports/sales/daily", reportHandler.DailySales)
	api.Get("/reports/top-selling", reportHandler.TopSelling)

	api.Post("/auth/register", authHandler.Register)
	api.Post("/auth/login", authHandler.Login)

	protected := api.Group("/", middleware.Protected())

	protected.Post("/items", itemHandler.Create)
	protected.Post("/inbounds", inboundHandler.Create)
	protected.Post("/sales", salesHandler.Create)

	return app
}
