package handler

import (
	domain "toki/internal/domain/inbound"
	uc "toki/internal/usecase/inbound"

	"github.com/gofiber/fiber/v2"
)

type InboundHandler struct {
	uc uc.Usecase
}

func NewInboundHandler(uc uc.Usecase) *InboundHandler {
	return &InboundHandler{uc}
}

func (h *InboundHandler) Create(c *fiber.Ctx) error {

	var req domain.Inbound

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	err := h.uc.Create(c.Context(), &req)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "inbound created",
	})
}
