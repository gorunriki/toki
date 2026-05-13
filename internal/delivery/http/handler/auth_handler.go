package handler

import (
	domain "toki/internal/domain/user"
	uc "toki/internal/usecase/auth"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	uc uc.Usecase
}

func NewAuthHandler(uc uc.Usecase) *AuthHandler {
	return &AuthHandler{uc}
}

func (h *AuthHandler) Register(c *fiber.Ctx) error {

	var req struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
		Role     string `json:"role"`
	}

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	err := h.uc.Register(c.Context(), &domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
		Role:     req.Role,
	})

	if err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "register success",
	})
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {

	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.BodyParser(&req); err != nil {
		return err
	}

	token, err := h.uc.Login(
		c.Context(),
		req.Email,
		req.Password,
	)

	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}
