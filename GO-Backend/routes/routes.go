package routes

import (
	"go-fiber-app/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, userHandler *handler.UserHandler, phoneHandler *handler.PhoneHandler) {
	// User routes
	r.Get("/users", userHandler.GetAllUsers)
	r.Post("/users", userHandler.CreateUser)
	r.Get("/users/:id", userHandler.GetUser)
	r.Get("/users/:id/details", userHandler.GetUserWithPhones) // Get user with all phone details
	r.Put("/users/:id", userHandler.UpdateUser)
	r.Delete("/users/:id", userHandler.DeleteUser)

	// Phone routes
	r.Post("/users/:userId/phones", phoneHandler.CreatePhone)
	r.Get("/users/:userId/phones", phoneHandler.GetPhonesByUser)    // Get all phones for a user
	r.Get("/users/:userId/phones/:id", phoneHandler.GetPhone)       // Get a specific phone
	r.Put("/users/:userId/phones/:id", phoneHandler.UpdatePhone)    // Update a specific phone
	r.Delete("/users/:userId/phones/:id", phoneHandler.DeletePhone) // Delete a specific phone
}
