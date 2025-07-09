package handler

import (
	"fmt"
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
	userID := c.Params("id")
	fmt.Printf("Received request to create phone for user ID: %s\n", userID)

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		fmt.Printf("Error converting user ID to ObjectID: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	var phone model.PhoneNumber
	if err := c.BodyParser(&phone); err != nil {
		fmt.Printf("Error parsing request body: %v\n", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	phone.UserID = userObjectID
	fmt.Printf("Creating phone: %+v\n", phone)

	if err := h.phoneService.CreatePhone(&phone); err != nil {
		fmt.Printf("Error creating phone in service: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	fmt.Printf("Phone created successfully: %+v\n", phone)
	return c.Status(fiber.StatusCreated).JSON(phone)
}

func (h *PhoneHandler) GetPhonesByUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	phones, err := h.phoneService.GetPhonesByUser(userObjectID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(phones)
}

func (h *PhoneHandler) UpdatePhone(c *fiber.Ctx) error {
	userID := c.Params("id")
	phoneID := c.Params("phoneId")

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	phoneObjectID, err := primitive.ObjectIDFromHex(phoneID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid phone ID"})
	}

	var phone model.PhoneNumber
	if err := c.BodyParser(&phone); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	phone.ID = phoneObjectID
	phone.UserID = userObjectID

	if err := h.phoneService.UpdatePhone(&phone); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(phone)
}

func (h *PhoneHandler) DeletePhone(c *fiber.Ctx) error {
	userID := c.Params("id")
	phoneID := c.Params("phoneId")

	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid user ID"})
	}

	phoneObjectID, err := primitive.ObjectIDFromHex(phoneID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid phone ID"})
	}

	if err := h.phoneService.DeletePhone(userObjectID, phoneObjectID); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusNoContent).Send(nil)
}
