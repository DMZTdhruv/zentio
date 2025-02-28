package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/personal-project/zentio/internal/config"
	authHandler "github.com/personal-project/zentio/internal/handlers"
	authRepo "github.com/personal-project/zentio/internal/repository"
	"github.com/personal-project/zentio/internal/routes"
	authService "github.com/personal-project/zentio/internal/services"
)

func SetUpRoutes(app fiber.Router) {
	setUpAuthRoutes(app)
}

func setUpAuthRoutes(app fiber.Router) {
	authRepository := authRepo.NewAuthRepository(config.GetDb())
	authServices := authService.NewAuthService(authRepository)
	authHandler := authHandler.NewAuthHandler(authServices)

	routes.AuthRoutes(app, authHandler)
}
