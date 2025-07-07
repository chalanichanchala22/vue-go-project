package main

import (
	"go-fiber-app/config"
	"go-fiber-app/handler"
	"go-fiber-app/repository"
	"go-fiber-app/routes"
	"go-fiber-app/service"
	"go-fiber-app/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	utils.LoadEnv()
	config.ConnectDB()

	app := fiber.New()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Allow  Vue app origin
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	// Root route - Health check/Welcome message
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Go Backend API",
			"status":  "running",
		})
	})

	db := config.GetDatabase()
	userRepo := repository.NewUserRepository(db)
	phoneRepo := repository.NewPhoneRepository(db)
	userService := service.NewUserService(userRepo, phoneRepo)
	phoneService := service.NewPhoneService(phoneRepo, userRepo)
	userHandler := handler.NewUserHandler(userService)
	phoneHandler := handler.NewPhoneHandler(phoneService)

	// API routes with prefix
	api := app.Group("/api")
	routes.RegisterUserRoutes(api, userHandler)
	routes.RegisterPhoneRoutes(api, phoneHandler)

	app.Listen(":8080")
}
