package handler

import (
	model "go-fiber-app/models"
	"go-fiber-app/service"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PhoneHandler struct {
	phoneService *service.PhoneService
}

func NewPhoneHandler(phoneService *service.PhoneService) *PhoneHandler {
	return &PhoneHandler{phoneService: phoneService}
}

func (h *PhoneHandler) CreatePhone(c *fiber.Ctx) error {
	var phone model.PhoneNumber
	if err := c.BodyParser(&phone); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	userID, err := primitive.ObjectIDFromHex(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}
	phone.UserID = userID

	if err := h.phoneService.CreatePhone(&phone); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "Phone created successfully", "phone": phone})
}

// GetPhonesByUser handles GET /users/:userId/phones
func (h *PhoneHandler) GetPhonesByUser(c *fiber.Ctx) error {
	userID, err := primitive.ObjectIDFromHex(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	phones, err := h.phoneService.GetPhonesByUser(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(phones)
}

// GetPhone handles GET /users/:userId/phones/:id
func (h *PhoneHandler) GetPhone(c *fiber.Ctx) error {
	// TODO: Implement logic to retrieve a specific phone by userId and id
	return c.SendStatus(fiber.StatusNotImplemented)
}

// UpdatePhone updates a specific phone for a user.
func (h *PhoneHandler) UpdatePhone(c *fiber.Ctx) error {
	// TODO: Implement the update logic here
	return c.SendStatus(fiber.StatusNotImplemented)
}

// DeletePhone deletes a specific phone by ID for a user.
func (h *PhoneHandler) DeletePhone(c *fiber.Ctx) error {
	userID, err := primitive.ObjectIDFromHex(c.Params("userId"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	phoneID, err := primitive.ObjectIDFromHex(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid phone ID"})
	}

	if err := h.phoneService.DeletePhone(userID, phoneID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Phone deleted successfully"})
}
