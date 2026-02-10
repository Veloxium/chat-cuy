package handler

import (
	"strconv"

	"github.com/Gylmynnn/chat-cuy-be/internal/service"
	"github.com/Gylmynnn/chat-cuy-be/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type MessageHandler struct {
	messageService service.MessageService
}

func NewMessageHandler(messageService service.MessageService) *MessageHandler {
	return &MessageHandler{messageService: messageService}
}

func (h *MessageHandler) GetMessages(c *fiber.Ctx) error {
	roomIDStr := c.Params("roomId")
	roomID, err := strconv.Atoi(roomIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid room ID")
	}

	messages, err := h.messageService.GetMessages(c.Context(), roomID)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to fetch messages")
	}

	if messages == nil {
		return c.Status(fiber.StatusOK).JSON([]any{})
	}

	return c.Status(fiber.StatusOK).JSON(messages)
}
