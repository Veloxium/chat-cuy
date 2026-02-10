package handler

import (
	"github.com/Gylmynnn/chat-cuy-be/internal/model"
	"github.com/Gylmynnn/chat-cuy-be/internal/service"
	"github.com/Gylmynnn/chat-cuy-be/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type ContactHandler struct {
	contactService service.ContactService
}

func NewContactHandler(contactService service.ContactService) *ContactHandler {
	return &ContactHandler{contactService: contactService}
}

func (h *ContactHandler) AddContact(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	var req model.AddContactRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if req.ContactUsername == "" {
		return utils.SendError(c, fiber.StatusBadRequest, "Contact username is required")
	}

	if err := h.contactService.AddContact(c.Context(), userID, req.ContactUsername); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, err.Error())
	}

	return utils.SendSuccess(c, fiber.StatusCreated, "Contact added successfully", nil)
}

func (h *ContactHandler) GetContacts(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	contacts, err := h.contactService.GetContacts(c.Context(), userID)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to fetch contacts")
	}

	if contacts == nil {
		return c.Status(fiber.StatusOK).JSON([]interface{}{})
	}

	return c.Status(fiber.StatusOK).JSON(contacts)
}
