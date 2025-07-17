package main

import (
	"context"
	"fmt"
	"go-fiber-app/config"
	"go-fiber-app/handler"
	model "go-fiber-app/models"
	"go-fiber-app/repository"
	"go-fiber-app/routes"
	"go-fiber-app/service"
	"go-fiber-app/utils"
	"log"
	"os"
	"time"

	_ "go-fiber-app/docs" // important for swag docs

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	jwtware "github.com/gofiber/jwt/v3"
	fiberSwagger "github.com/swaggo/fiber-swagger" // fiber swagger middleware
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

	// Seed default data
	seedData(userRepo)

	userService := service.NewUserService(userRepo)
	userService.SetPhoneRepository(phoneRepo)
	userHandler := handler.NewUserHandler(userService)

	phoneService := service.NewPhoneService(phoneRepo)
	phoneHandler := handler.NewPhoneHandler(phoneService)

	authHandler := handler.NewAuthHandler(userService)

	// JWT middleware for protected routes
	app.Use("/api/users", jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))

	routes.RegisterRoutes(app, userHandler, phoneHandler, authHandler)

	fmt.Println("Server starting on :8080...")
	log.Fatal(app.Listen(":8080"))
}

// seedData creates default users if they don't exist
func seedData(userRepo *repository.UserRepository) {
	ctx := context.Background()

	// Check if admin user already exists
	_, err := userRepo.FindUserByEmail(ctx, "admin@example.com")
	if err == nil {
		fmt.Println("Seed data: Admin user already exists, skipping...")
		return
	}

	// Only create if user doesn't exist (err should be mongo.ErrNoDocuments)
	if err != mongo.ErrNoDocuments {
		fmt.Printf("Seed data: Error checking for existing admin user: %v\n", err)
		return
	}

	fmt.Println("Seed data: Creating default admin user...")

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("password123"), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("Seed data: Error hashing password: %v\n", err)
		return
	}

	// Create admin user
	adminUser := &model.User{
		Name:     "Admin User",
		Email:    "admin@example.com",
		Password: string(hashedPassword),
		NIC:      "123456789V",
		Address:  "123 Admin Street",
		Birthday: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Gender:   "Other",
		Photo:    "",
	}

	if err := userRepo.CreateUser(ctx, adminUser); err != nil {
		fmt.Printf("Seed data: Error creating admin user: %v\n", err)
		return
	}

	fmt.Printf("Seed data: Admin user created successfully with ID: %s\n", adminUser.ID.Hex())
	fmt.Println("Seed data: Login credentials - Email: admin@example.com, Password: password123")
}
