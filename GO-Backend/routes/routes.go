package routes

import (
	"go-fiber-app/handler"

	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App, userHandler *handler.UserHandler, phoneHandler *handler.PhoneHandler, authHandler *handler.AuthHandler) {
	api := app.Group("/api") // Group everything under /api

	// === Public Routes ===
	api.Post("/auth/login", authHandler.Login)
	api.Post("/auth/register", userHandler.CreateUser) // Allow public registration
	api.Get("/users/list", userHandler.GetAllUsers)    // Temporary: List users for debugging

	// === Protected Routes ===
	userGroup := api.Group("/users")

	// User routes
	userGroup.Get("/", userHandler.GetAllUsers)
	userGroup.Get("/with-phones", userHandler.GetAllUsersWithPhones)
	userGroup.Post("/", userHandler.CreateUser)
	userGroup.Post("/set-password", userHandler.SetPasswordForExistingUser) // Temporary migration endpoint
	userGroup.Get("/:id", userHandler.GetUser)
	userGroup.Get("/:id/details", userHandler.GetUser)
	userGroup.Put("/:id", userHandler.UpdateUser)
	userGroup.Put("/:id/password", userHandler.UpdateUserPassword)
	userGroup.Delete("/:id", userHandler.DeleteUser)
	userGroup.Get("/:id/with-phones", userHandler.GetUserWithPhones)

	// Phone routes
	userGroup.Get("/:id/phones", phoneHandler.GetPhonesByUser)
	userGroup.Post("/:id/phones", phoneHandler.CreatePhone)
	userGroup.Put("/:id/phones/:phoneId", phoneHandler.UpdatePhone)
	userGroup.Delete("/:id/phones/:phoneId", phoneHandler.DeletePhone)

	// Test route for file upload debugging
	api.Post("/test-upload", func(c *fiber.Ctx) error {
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
