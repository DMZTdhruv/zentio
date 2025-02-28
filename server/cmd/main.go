package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/personal-project/zentio/cmd/api"
	"github.com/personal-project/zentio/internal/config"

	"github.com/personal-project/zentio/internal/utils"
)

func main() {
	config.LoadEnv()
	utils.InitJwtSecret()
	config.ConnectToDb()

	var port string
	port, err := config.GetEnv("PORT")
	if err != nil {
		fmt.Printf("Port is not defined")
		port = "3001"
	}

	app := fiber.New()

	// api routes for auth
	apiGroup := app.Group("/api")
	api.SetUpRoutes(apiGroup)

	app.Listen(":" + port)
}
