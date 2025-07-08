package main

import (
	"fmt"
	"go-fiber-app/config"
	"go-fiber-app/handler"
	"go-fiber-app/repository"
	"go-fiber-app/routes"
	"go-fiber-app/service"
	"go-fiber-app/utils"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	fmt.Println("Starting Go Backend Server...")

	utils.LoadEnv()
	config.ConnectDB()

	app := fiber.New()

	// CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173", // Allow Vue app origin
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization, X-Requested-With",
		AllowCredentials: true,
	}))

	// Serve static files for uploads
	app.Static("/uploads", "./storage/uploads")

	// Root route - Health check/Welcome message
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"message": "Welcome to Go Backend API",
			"status":  "running",
		})
	})

	db := config.GetDatabase()
	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	// API routes with prefix
	api := app.Group("/api")
	routes.RegisterRoutes(api, userHandler)

	fmt.Println("Server starting on :8080...")
	log.Fatal(app.Listen(":8080"))
}
