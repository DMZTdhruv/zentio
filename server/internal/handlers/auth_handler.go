package auth

import (
	"github.com/gofiber/fiber/v2"
	"github.com/personal-project/zentio/internal/schema"
	"github.com/personal-project/zentio/internal/services"
	"github.com/personal-project/zentio/internal/zentio"
)

// constructor for the auth service?
type AuthHandler struct {
	Service *auth.AuthService
}

func NewAuthHandler(serivce *auth.AuthService) *AuthHandler {
	return &AuthHandler{
		Service: serivce,
	}
}

func (h *AuthHandler) SignUpHandler(c *fiber.Ctx) error {
	var user schema.User
	if err := c.BodyParser(&user); err != nil {
		return zentio.ErrorResponse(c, fiber.StatusAccepted, err.Error(), nil)
	}

	token, err := h.Service.SignUpUser(&user)
	if err != nil {
		return zentio.ErrorResponse(c, fiber.StatusAccepted, err.Error(), nil)
	}

	cookie := fiber.Cookie{
		Name:     "auth",
		Value:    token,
		Path:     "/",
		HTTPOnly: true,
		Secure:   true,
		SameSite: "Strict",
		MaxAge:   86400,
	}
	c.Cookie(&cookie)
	return zentio.SuccessResponse(c, fiber.StatusAccepted, "successfully created user", nil)
}
