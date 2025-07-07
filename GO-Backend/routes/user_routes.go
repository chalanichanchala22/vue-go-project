package routes

import (
	"go-fiber-app/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterUserRoutes(r fiber.Router, h *handler.UserHandler) {
	r.Get("/users", h.GetAllUsers)
	r.Post("/users", h.CreateUser)
	r.Get("/users/:id", h.GetUser)
	r.Put("/users/:id", h.UpdateUser)
	r.Delete("/users/:id", h.DeleteUser)

}
