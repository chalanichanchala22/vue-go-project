package routes

import (
	"go-fiber-app/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterPhoneRoutes(r fiber.Router, h *handler.PhoneHandler) {
	r.Post("/users/:userId/phones", h.CreatePhone)
	r.Get("/users/:userId/phones", h.GetPhonesByUser)    // Get all phones for a user
	r.Get("/users/:userId/phones/:id", h.GetPhone)       // Get a specific phone
	r.Put("/users/:userId/phones/:id", h.UpdatePhone)    // Update a specific phone
	r.Delete("/users/:userId/phones/:id", h.DeletePhone) // Delete a specific phone
}
