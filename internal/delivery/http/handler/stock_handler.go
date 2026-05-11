package handler

import (
	"strconv"

	"toki/internal/usecase/stock"

	"github.com/gofiber/fiber/v2"
)

type StockHandler struct {
	uc stock.Usecase
}

func NewStockHandler(uc stock.Usecase) *StockHandler {
	return &StockHandler{uc}
}

func (h *StockHandler) GetStock(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("item_id"))

	qty, err := h.uc.GetStock(c.Context(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"item_id": id,
		"stock":   qty,
	})
}
