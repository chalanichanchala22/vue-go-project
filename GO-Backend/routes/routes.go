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
	r.Get("/users/:id/details", userHandler.GetUser) 
	r.Put("/users/:id", userHandler.UpdateUser)
	r.Delete("/users/:id", userHandler.DeleteUser)

}
