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

	_ "go-fiber-app/docs" // important for swag docs

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	fiberSwagger "github.com/swaggo/fiber-swagger" // fiber swagger middleware
)

// @title           Go Fiber User API
// @version         1.0
// @description     A simple user management API using Fiber and MongoDB.
// @host            localhost:8080
// @BasePath        /api

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

	// Swagger route
	app.Get("/swagger/*", fiberSwagger.WrapHandler)

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

	userService := service.NewUserService(userRepo)
	userService.SetPhoneRepository(phoneRepo)
	userHandler := handler.NewUserHandler(userService)

	phoneService := service.NewPhoneService(phoneRepo)
	phoneHandler := handler.NewPhoneHandler(phoneService)

	// API routes with prefix
	api := app.Group("/api")
	routes.RegisterRoutes(api, userHandler, phoneHandler)

	fmt.Println("Server starting on :8080...")
	log.Fatal(app.Listen(":8080"))
}
