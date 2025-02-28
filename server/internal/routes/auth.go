package routes

import (
	"github.com/gofiber/fiber/v2"
	auth "github.com/personal-project/zentio/internal/handlers"
)

func AuthRoutes(app fiber.Router, handler *auth.AuthHandler) {
	authGroup := app.Group("/auth")
	authGroup.Post("/sign-up", handler.SignUpHandler)
}
