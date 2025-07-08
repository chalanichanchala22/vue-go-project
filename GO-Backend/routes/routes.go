package routes

import (
	"go-fiber-app/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, userHandler *handler.UserHandler) {
	// User routes
	r.Get("/users", userHandler.GetAllUsers)
	r.Post("/users", userHandler.CreateUser)
	r.Get("/users/:id", userHandler.GetUser)
	r.Get("/users/:id/details", userHandler.GetUser)
	r.Put("/users/:id", userHandler.UpdateUser)
	r.Delete("/users/:id", userHandler.DeleteUser)

	// Test route for file upload debugging
	r.Post("/test-upload", func(c *fiber.Ctx) error {
		form, err := c.MultipartForm()
		if err != nil {
			return c.JSON(fiber.Map{"error": "Not multipart form", "details": err.Error()})
		}

		files := form.File["photo"]
		if len(files) > 0 {
			return c.JSON(fiber.Map{
				"message":  "File received successfully",
				"filename": files[0].Filename,
				"size":     files[0].Size,
			})
		}

		return c.JSON(fiber.Map{"message": "No file received", "form_fields": form.Value})
	})
}
