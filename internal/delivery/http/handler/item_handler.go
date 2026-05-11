package handler

import (
	"strconv"
	domain "toki/internal/domain/item"
	uc "toki/internal/usecase/item"

	"github.com/gofiber/fiber/v2"
)

type ItemHandler struct {
	uc uc.Usecase
}

func NewItemHandler(uc uc.Usecase) *ItemHandler {
	return &ItemHandler{uc}
}

func (h *ItemHandler) Create(c *fiber.Ctx) error {
	var req struct {
		Name      string  `json:"name"`
		SKU       string  `json:"sku"`
		PriceSell float64 `json:"price_sell"`
		PriceBuy  float64 `json:"price_buy"`
	}

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "invalid request",
		})
	}

	err := h.uc.Create(c.Context(), &domain.Item{
		Name:      req.Name,
		SKU:       req.SKU,
		PriceSell: req.PriceSell,
		PriceBuy:  req.PriceBuy,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "item created",
	})
}

func (h *ItemHandler) FindAll(c *fiber.Ctx) error {
	items, err := h.uc.FindAll(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(items)
}

func (h *ItemHandler) FindByID(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	item, err := h.uc.FindByID(c.Context(), id)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "not found"})
	}

	return c.JSON(item)
}

func (h *ItemHandler) Update(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	var req struct {
		Name      string  `json:"name"`
		PriceSell float64 `json:"price_sell"`
	}

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	err := h.uc.Update(c.Context(), &domain.Item{
		ID:        id,
		Name:      req.Name,
		PriceSell: req.PriceSell,
	})
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "updated"})
}

func (h *ItemHandler) Delete(c *fiber.Ctx) error {
	id, _ := strconv.Atoi(c.Params("id"))

	err := h.uc.Delete(c.Context(), id)
	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{"message": "deleted"})
}
