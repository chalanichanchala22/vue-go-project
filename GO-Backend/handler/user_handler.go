package handler

//used to handle HTTP requests.
import (
	"fmt"
	model "go-fiber-app/models"
	"go-fiber-app/service"
	"os"
	"path/filepath"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
	//primitive is from MongoDB, used to convert string IDs to MongoDB’s ObjectID format.
)

type UserHandler struct { //a struct to group all user-related route functions.
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	// Parse multipart form
	form, err := c.MultipartForm()
	if err != nil {
		// Fallback to JSON parsing if not multipart
		var user model.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		// Ensure photo field is initialized even for JSON requests
		if user.Photo == "" {
			user.Photo = ""
		}
		fmt.Printf("JSON request - User data: %+v\n", user)
		if err := h.userService.CreateUser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.Status(fiber.StatusCreated).JSON(user)
	}

	// Extract form fields
	var user model.User

	if name := form.Value["name"]; len(name) > 0 {
		user.Name = name[0]
	}
	if email := form.Value["email"]; len(email) > 0 {
		user.Email = email[0]
	}
	if nic := form.Value["nic"]; len(nic) > 0 {
		user.NIC = nic[0]
	}
	if address := form.Value["address"]; len(address) > 0 {
		user.Address = address[0]
	}
	if birthday := form.Value["birthday"]; len(birthday) > 0 {
		if parsedTime, err := time.Parse("2006-01-02T15:04:05.000Z", birthday[0]); err == nil {
			user.Birthday = parsedTime
		}
	}
	if gender := form.Value["gender"]; len(gender) > 0 {
		user.Gender = gender[0]
	}

	// Initialize photo field with default value
	user.Photo = ""

	// Handle file upload
	files := form.File["photo"]
	if len(files) > 0 {
		fmt.Printf("Photo file found: %s\n", files[0].Filename)
		file := files[0]

		// Validate file type
		ext := filepath.Ext(file.Filename)
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Only image files are allowed"})
		}

		// Create uploads directory if it doesn't exist
		uploadDir := "./storage/uploads"
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create upload directory"})
		}

		// Generate unique filename
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		filePath := fmt.Sprintf("%s/%s", uploadDir, filename)

		// Save file
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not save file"})
		}

		// Store relative path in user
		user.Photo = fmt.Sprintf("/uploads/%s", filename)
		fmt.Printf("Photo path set to: %s\n", user.Photo)
	} else {
		fmt.Printf("No photo file found in form\n")
	}

	fmt.Printf("User data before saving: %+v\n", user)
	if err := h.userService.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	fmt.Printf("User saved successfully: %+v\n", user)
	return c.Status(fiber.StatusCreated).JSON(user)
}

func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	user, err := h.userService.GetUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(user)
}

func (h *UserHandler) GetUserWithPhones(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}
	userWithPhones, err := h.userService.GetUserWithPhones(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.JSON(userWithPhones)
}

func (h *UserHandler) GetAllUsersWithPhones(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsersWithPhones()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	// Try to parse as multipart form first
	form, err := c.MultipartForm()
	if err != nil {
		// Fallback to JSON parsing if not multipart
		var user model.User
		if err := c.BodyParser(&user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
		user.ID = userID

		if err := h.userService.UpdateUser(&user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}
		return c.JSON(user)
	}

	// Get existing user first
	existingUser, err := h.userService.GetUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Update fields from form
	var user model.User
	user.ID = userID

	if name := form.Value["name"]; len(name) > 0 {
		user.Name = name[0]
	} else {
		user.Name = existingUser.Name
	}
	if email := form.Value["email"]; len(email) > 0 {
		user.Email = email[0]
	} else {
		user.Email = existingUser.Email
	}
	if nic := form.Value["nic"]; len(nic) > 0 {
		user.NIC = nic[0]
	} else {
		user.NIC = existingUser.NIC
	}
	if address := form.Value["address"]; len(address) > 0 {
		user.Address = address[0]
	} else {
		user.Address = existingUser.Address
	}
	if birthday := form.Value["birthday"]; len(birthday) > 0 {
		if parsedTime, err := time.Parse("2006-01-02T15:04:05.000Z", birthday[0]); err == nil {
			user.Birthday = parsedTime
		}
	} else {
		user.Birthday = existingUser.Birthday
	}
	if gender := form.Value["gender"]; len(gender) > 0 {
		user.Gender = gender[0]
	} else {
		user.Gender = existingUser.Gender
	}

	// Keep existing photo by default
	user.Photo = existingUser.Photo

	// Handle file upload if provided
	files := form.File["photo"]
	if len(files) > 0 {
		file := files[0]

		// Validate file type
		ext := filepath.Ext(file.Filename)
		if ext != ".jpg" && ext != ".jpeg" && ext != ".png" && ext != ".gif" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Only image files are allowed"})
		}

		// Create uploads directory if it doesn't exist
		uploadDir := "./storage/uploads"
		if err := os.MkdirAll(uploadDir, 0755); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create upload directory"})
		}

		// Generate unique filename
		filename := fmt.Sprintf("%d_%s", time.Now().Unix(), file.Filename)
		filePath := fmt.Sprintf("%s/%s", uploadDir, filename)

		// Save file
		if err := c.SaveFile(file, filePath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not save file"})
		}

		// Update photo path
		user.Photo = fmt.Sprintf("/uploads/%s", filename)
	}

	if err := h.userService.UpdateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(user)
}

func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	if err := h.userService.DeleteUser(userID); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}
	return c.SendStatus(fiber.StatusNoContent)
}
