package handler

import (
	domain "toki/internal/domain/sales"
	uc "toki/internal/usecase/sales"

	"github.com/gofiber/fiber/v2"
)

type SalesHandler struct {
	uc uc.Usecase
}

func NewSalesHandler(uc uc.Usecase) *SalesHandler {
	return &SalesHandler{uc}
}

func (h *SalesHandler) Create(c *fiber.Ctx) error {

	var req domain.Sales

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	err := h.uc.Create(c.Context(), &req)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "sales created",
	})
}
