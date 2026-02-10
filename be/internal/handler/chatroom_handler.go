package handler

import (
	"strconv"

	"github.com/Gylmynnn/chat-cuy-be/internal/service"
	"github.com/Gylmynnn/chat-cuy-be/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type ChatRoomHandler struct {
	chatRoomService service.ChatRoomService
}

func NewChatRoomHandler(chatRoomService service.ChatRoomService) *ChatRoomHandler {
	return &ChatRoomHandler{chatRoomService: chatRoomService}
}

func (h *ChatRoomHandler) GetUserChatRooms(c *fiber.Ctx) error {
	userIDStr := c.Params("userId")
	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid user ID")
	}

	rooms, err := h.chatRoomService.GetUserChatRooms(c.Context(), userID)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to fetch chat rooms")
	}

	if rooms == nil {
		return c.Status(fiber.StatusOK).JSON([]any{})
	}

	return c.Status(fiber.StatusOK).JSON(rooms)
}

func (h *ChatRoomHandler) CreatePrivateRoom(c *fiber.Ctx) error {
	type request struct {
		UserID1 int `json:"user_id_1"`
		UserID2 int `json:"user_id_2"`
	}

	var req request
	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	room, err := h.chatRoomService.GetOrCreatePrivateRoom(c.Context(), req.UserID1, req.UserID2)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to create chat room")
	}

	return utils.SendSuccess(c, fiber.StatusOK, "Chat room ready", room)
}
