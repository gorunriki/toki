package handler

import (
	"strconv"

	"toki/internal/domain/item"
	itemUC "toki/internal/usecase/item"

	"github.com/gofiber/fiber/v2"
)

type ItemHandler struct {
	uc itemUC.Usecase
}

func NewItemHandler(
	uc itemUC.Usecase,
) *ItemHandler {

	return &ItemHandler{
		uc: uc,
	}
}

func (h *ItemHandler) Create(
	c *fiber.Ctx,
) error {

	var req item.Item

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	err := h.uc.Create(
		c.Context(),
		&req,
	)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		fiber.Map{
			"message": "item created",
		},
	)
}

func (h *ItemHandler) FindAll(
	c *fiber.Ctx,
) error {

	items, err := h.uc.FindAll(
		c.Context(),
	)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(items)
}

func (h *ItemHandler) FindByID(
	c *fiber.Ctx,
) error {

	id, err := strconv.Atoi(
		c.Params("id"),
	)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid id",
			},
		)
	}

	it, err := h.uc.FindByID(
		c.Context(),
		id,
	)

	if err != nil {
		return c.Status(404).JSON(
			fiber.Map{
				"error": "item not found",
			},
		)
	}

	return c.JSON(it)
}

func (h *ItemHandler) Update(
	c *fiber.Ctx,
) error {

	id, err := strconv.Atoi(
		c.Params("id"),
	)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid id",
			},
		)
	}

	var req item.Item

	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	req.ID = id

	err = h.uc.Update(
		c.Context(),
		&req,
	)

	if err != nil {
		return c.Status(500).JSON(
			fiber.Map{
				"error": err.Error(),
			},
		)
	}

	return c.JSON(
		fiber.Map{
			"message": "item updated",
		},
	)
}

func (h *ItemHandler) Delete(
	c *fiber.Ctx,
) error {

	id, err := strconv.Atoi(
		c.Params("id"),
	)

	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"error": "invalid id",
			},
		)
	}

	err = h.uc.Delete(
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

	return c.JSON(
		fiber.Map{
			"message": "item deleted",
		},
	)
}
