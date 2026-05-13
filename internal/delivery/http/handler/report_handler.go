package handler

import (
	uc "toki/internal/usecase/report"

	"github.com/gofiber/fiber/v2"
)

type ReportHandler struct {
	uc uc.Usecase
}

func NewReportHandler(uc uc.Usecase) *ReportHandler {
	return &ReportHandler{uc}
}

func (h *ReportHandler) Stocks(c *fiber.Ctx) error {

	data, err := h.uc.GetStocks(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(data)
}

func (h *ReportHandler) DailySales(c *fiber.Ctx) error {

	data, err := h.uc.GetDailySales(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(data)
}

func (h *ReportHandler) TopSelling(c *fiber.Ctx) error {

	data, err := h.uc.GetTopSelling(c.Context())
	if err != nil {
		return err
	}

	return c.JSON(data)
}
