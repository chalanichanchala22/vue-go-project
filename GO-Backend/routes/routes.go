package routes

import (
	"go-fiber-app/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(r fiber.Router, userHandler *handler.UserHandler, phoneHandler *handler.PhoneHandler) {
	// User routes
	r.Get("/users", userHandler.GetAllUsers)
	r.Get("/users/with-phones", userHandler.GetAllUsersWithPhones)
	r.Post("/users", userHandler.CreateUser)
	r.Get("/users/:id", userHandler.GetUser)
	r.Get("/users/:id/details", userHandler.GetUser)
	r.Put("/users/:id", userHandler.UpdateUser)
	r.Delete("/users/:id", userHandler.DeleteUser)

	// Phone routes
	r.Get("/users/:id/phones", phoneHandler.GetPhonesByUser)
	r.Post("/users/:id/phones", phoneHandler.CreatePhone)
	r.Put("/users/:id/phones/:phoneId", phoneHandler.UpdatePhone)
	r.Delete("/users/:id/phones/:phoneId", phoneHandler.DeletePhone)

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
