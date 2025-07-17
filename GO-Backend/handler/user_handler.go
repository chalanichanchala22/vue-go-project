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
	"golang.org/x/crypto/bcrypt"
	//primitive is from MongoDB, used to convert string IDs to MongoDB’s ObjectID format.
)

type UpdatePasswordRequest struct {
	CurrentPassword string `json:"currentPassword" validate:"required" example:"oldpassword123"`
	NewPassword     string `json:"newPassword" validate:"required,min=6" example:"newpassword123"`
	ConfirmPassword string `json:"confirmPassword" validate:"required" example:"newpassword123"`
}

type CreateUserWithPasswordRequest struct {
	Name            string `json:"name" validate:"required" example:"John Doe"`
	Email           string `json:"email" validate:"required,email" example:"john.doe@example.com"`
	NIC             string `json:"nic" validate:"required" example:"123456789V"`
	Address         string `json:"address" validate:"required" example:"123 Main St, City"`
	Birthday        string `json:"birthday" validate:"required" example:"1990-01-15" format:"date"`
	Gender          string `json:"gender" validate:"required" example:"Male"`
	Password        string `json:"password" validate:"required,min=6" example:"password123"`
	ConfirmPassword string `json:"confirmPassword" validate:"required" example:"password123"` // Only for validation
}

type UpdateUserRequest struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	NIC      string `json:"nic,omitempty"`
	Address  string `json:"address,omitempty"`
	Birthday string `json:"birthday,omitempty"` // Handle as string for parsing
	Gender   string `json:"gender,omitempty"`
}

// a struct to group all user-related route functions.
type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService: userService}
}

// CreateUser godoc
// @Summary      Create a new user with password
// @Description  Create a new user account with email, password, and confirm password validation
// @Tags         Users
// @Accept       json,multipart/form-data
// @Produce      json
// @Param        request body CreateUserWithPasswordRequest true "User creation data with password (JSON)"
// @Param        name            formData string false "User's full name"
// @Param        email           formData string false "User's email address"
// @Param        nic             formData string false "National ID number"
// @Param        address         formData string false "User's address"
// @Param        birthday        formData string false "Birthday in YYYY-MM-DD format (simple date)"
// @Param        gender          formData string false "Gender (Male/Female)"
// @Param        password        formData string false "Password (minimum 6 characters)"
// @Param        confirmPassword formData string false "Confirm password (must match password)"
// @Param        photo           formData file   false "User's profile image (jpg/png/gif)"
// @Success      201  {object}  model.User
// @Failure      400  {object}  map[string]string  "Invalid request or validation errors"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /users [post]
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	// Try to parse as multipart form first
	form, err := c.MultipartForm()
	if err != nil {
		// Fallback to JSON parsing if not multipart
		var req CreateUserWithPasswordRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Validate that password and confirm password match
		if req.Password != req.ConfirmPassword {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password and confirm password do not match"})
		}

		// Validate password length
		if len(req.Password) < 6 {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password must be at least 6 characters long"})
		}

		// Hash password
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
		}

		// Parse birthday
		birthday, err := time.Parse("2006-01-02", req.Birthday)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid birthday format. Use YYYY-MM-DD"})
		}

		// Create user object
		user := &model.User{
			Name:     req.Name,
			Email:    req.Email,
			NIC:      req.NIC,
			Address:  req.Address,
			Birthday: birthday,
			Gender:   req.Gender,
			Password: string(hashedPassword),
			Photo:    "", // Default empty photo
		}

		// Save user
		if err := h.userService.CreateUser(user); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(fiber.StatusCreated).JSON(user)
	}

	// Handle multipart form data
	var user model.User

	// Extract and validate required fields from form
	if name := form.Value["name"]; len(name) > 0 {
		user.Name = name[0]
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Name is required"})
	}

	if email := form.Value["email"]; len(email) > 0 {
		user.Email = email[0]
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Email is required"})
	}

	if nic := form.Value["nic"]; len(nic) > 0 {
		user.NIC = nic[0]
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "NIC is required"})
	}

	if address := form.Value["address"]; len(address) > 0 {
		user.Address = address[0]
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Address is required"})
	}

	if gender := form.Value["gender"]; len(gender) > 0 {
		user.Gender = gender[0]
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Gender is required"})
	}

	// Handle password validation
	var password, confirmPassword string
	if pwd := form.Value["password"]; len(pwd) > 0 {
		password = pwd[0]
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password is required"})
	}

	if confPwd := form.Value["confirmPassword"]; len(confPwd) > 0 {
		confirmPassword = confPwd[0]
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Confirm password is required"})
	}

	// Validate that password and confirm password match
	if password != confirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password and confirm password do not match"})
	}

	// Validate password length
	if len(password) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Password must be at least 6 characters long"})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	user.Password = string(hashedPassword)

	// Handle birthday
	if birthday := form.Value["birthday"]; len(birthday) > 0 {
		birthdayStr := birthday[0]

		// Parse only simple date format (YYYY-MM-DD)
		parsedTime, err := time.Parse("2006-01-02", birthdayStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid birthday format. Expected YYYY-MM-DD format"})
		}

		user.Birthday = parsedTime
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Birthday is required"})
	}

	// Handle file upload if provided
	user.Photo = "" // Default empty photo
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
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create upload directory"})
		}

		// Generate unique filename
		timestamp := time.Now().Unix()
		filename := fmt.Sprintf("%d_%s", timestamp, file.Filename)
		filepath := fmt.Sprintf("%s/%s", uploadDir, filename)

		// Save the file
		if err := c.SaveFile(file, filepath); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to save file"})
		}

		// Store relative path in user object
		user.Photo = fmt.Sprintf("/uploads/%s", filename)
	}

	// Save user
	if err := h.userService.CreateUser(&user); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(user)
}

// GetAllUsers godoc
// @Summary      Get all users
// @Description  Retrieve a list of all users
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.User
// @Failure      500  {object}  map[string]string
// @Router       /users [get]
func (h *UserHandler) GetAllUsers(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsers()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

// GetUser godoc
// @Summary      Get a user by ID
// @Description  Retrieve a specific user by their ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  model.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [get]
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

// GetUserWithPhones godoc
// @Summary      Get a user with phone numbers
// @Description  Retrieve a specific user with all their phone numbers
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      200  {object}  model.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id}/with-phones [get]
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

// GetAllUsersWithPhones godoc
// @Summary      Get all users with their phone numbers
// @Description  Retrieve all users along with their associated phone numbers
// @Tags         Users
// @Accept       json
// @Produce      json
// @Success      200  {array}   model.User
// @Failure      500  {object}  map[string]string
// @Router       /users/with-phones [get]
func (h *UserHandler) GetAllUsersWithPhones(c *fiber.Ctx) error {
	users, err := h.userService.GetAllUsersWithPhones()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}

// UpdateUser godoc
// @Summary      Update a user
// @Description  Update an existing user with optional photo upload (multipart/form-data) or JSON data
// @Tags         Users
// @Accept       multipart/form-data,json
// @Produce      json
// @Param        id        path      string            true   "User ID"
// @Param        request   body      UpdateUserRequest false  "User update data (JSON)"
// @Param        name      formData  string            false  "User's full name"
// @Param        email     formData  string            false  "User's email address"
// @Param        nic       formData  string            false  "National ID number"
// @Param        address   formData  string            false  "User's address"
// @Param        birthday  formData  string            false  "Birthday in YYYY-MM-DD format (simple date)"
// @Param        gender    formData  string            false  "Gender (e.g. Male or Female)"
// @Param        photo     formData  file              false  "User's profile image (jpg/png/gif)"
// @Success      200  {object}  model.User
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Failure      500  {object}  map[string]string
// @Router       /users/{id} [put]
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
		var req UpdateUserRequest
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
		}

		// Get existing user first
		existingUser, err := h.userService.GetUser(userID)
		if err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}

		// Create updated user object, keeping existing values for fields not provided
		var user model.User
		user.ID = userID
		user.Password = existingUser.Password // Keep existing password

		if req.Name != "" {
			user.Name = req.Name
		} else {
			user.Name = existingUser.Name
		}

		if req.Email != "" {
			user.Email = req.Email
		} else {
			user.Email = existingUser.Email
		}

		if req.NIC != "" {
			user.NIC = req.NIC
		} else {
			user.NIC = existingUser.NIC
		}

		if req.Address != "" {
			user.Address = req.Address
		} else {
			user.Address = existingUser.Address
		}

		if req.Gender != "" {
			user.Gender = req.Gender
		} else {
			user.Gender = existingUser.Gender
		}

		// Handle birthday parsing
		if req.Birthday != "" {
			birthday, err := time.Parse("2006-01-02", req.Birthday)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid birthday format. Use YYYY-MM-DD"})
			}
			user.Birthday = birthday
		} else {
			user.Birthday = existingUser.Birthday
		}

		// Keep existing photo
		user.Photo = existingUser.Photo

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
		// Parse only simple date format (YYYY-MM-DD)
		if parsedTime, err := time.Parse("2006-01-02", birthday[0]); err == nil {
			user.Birthday = parsedTime
		} else {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid birthday format. Expected YYYY-MM-DD format"})
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

// DeleteUser godoc
// @Summary      Delete a user
// @Description  Delete a user by ID
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "User ID"
// @Success      204  "No Content"
// @Failure      400  {object}  map[string]string
// @Failure      404  {object}  map[string]string
// @Router       /users/{id} [delete]
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

// UpdateUserPassword godoc
// @Summary      Update user password
// @Description  Update a user's password with current password verification
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        id      path      string                 true  "User ID"
// @Param        request body      UpdatePasswordRequest  true  "Password update data"
// @Param        currentPassword  body  string  true  "Current password for verification"  example("oldpassword123")
// @Param        newPassword      body  string  true  "New password (minimum 6 characters)" example("newpassword123")
// @Param        confirmPassword  body  string  true  "Confirm new password (must match newPassword)" example("newpassword123")
// @Success      200  {object}  map[string]string  "Password updated successfully"
// @Failure      400  {object}  map[string]string  "Invalid request or validation errors"
// @Failure      401  {object}  map[string]string  "Current password is incorrect"
// @Failure      404  {object}  map[string]string  "User not found"
// @Failure      500  {object}  map[string]string  "Internal server error"
// @Router       /users/{id}/password [put]
func (h *UserHandler) UpdateUserPassword(c *fiber.Ctx) error {
	id := c.Params("id")
	userID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var req UpdatePasswordRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}

	// Validate that new password and confirm password match
	if req.NewPassword != req.ConfirmPassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "New password and confirm password do not match"})
	}

	// Validate password length
	if len(req.NewPassword) < 6 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "New password must be at least 6 characters long"})
	}

	// Get existing user
	existingUser, err := h.userService.GetUser(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
	}

	// Verify current password
	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(req.CurrentPassword)); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Current password is incorrect"})
	}

	// Hash new password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash new password"})
	}

	// Update user with new password
	existingUser.Password = string(hashedPassword)
	if err := h.userService.UpdateUser(existingUser); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to update password"})
	}

	return c.JSON(fiber.Map{"message": "Password updated successfully"})
}
