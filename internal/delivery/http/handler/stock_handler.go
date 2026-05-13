package handler

import (
	"strconv"

	stockUC "toki/internal/usecase/stock"

	"github.com/gofiber/fiber/v2"
)

type StockHandler struct {
	uc stockUC.Usecase
}

func NewStockHandler(
	uc stockUC.Usecase,
) *StockHandler {

	return &StockHandler{
		uc: uc,
	}
}

func (h *StockHandler) GetStock(
	c *fiber.Ctx,
) error {

	id, err := strconv.Atoi(
		c.Params("item_id"),
	)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid item id",
			},
		)
	}

	qty, err := h.uc.GetStock(
		c.Context(),
		id,
	)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(fiber.Map{
		"item_id": id,
		"stock":   qty,
	})
}
