package auth

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/personal-project/zentio/internal/schema"
	auth "github.com/personal-project/zentio/internal/services"
	"github.com/personal-project/zentio/internal/utils"
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

	err := h.Service.SignUpUser(&user)
	if err != nil {
		return zentio.ErrorResponse(c, fiber.ErrBadRequest.Code, err.Error(), nil)
	}

	// generating a jwt token
	token, err := utils.GenerateToken(user.ID, user.Username)
	if err != nil {
		return err
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

	response := schema.SignedUser{
		Username: user.Username,
		Name:     user.Name,
	}
	return zentio.SuccessResponse(c, fiber.StatusCreated, "successfully created user", response)
}

func (h *AuthHandler) SignInHandler(c *fiber.Ctx) error {
	var user schema.SignInUser
	fmt.Println("hello from the signin function")
	if err := c.BodyParser(&user); err != nil {
		return err
	}

	verifiedUser, err := h.Service.SignInUser(user.Email, user.Password)
	if err != nil {
		return zentio.ErrorResponse(c, fiber.StatusBadGateway, err.Error(), nil)
	}

	token, err := utils.GenerateToken(verifiedUser.ID, verifiedUser.Username)
	if err != nil {
		return err
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

	response := schema.SignedUser{
		Username: verifiedUser.Username,
		Name:     verifiedUser.Name,
	}
	return zentio.SuccessResponse(c, fiber.StatusCreated, "successfully signed user", response)
}
