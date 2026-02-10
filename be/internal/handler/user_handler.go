package handler

import (
	"fmt"
	"path/filepath"
	"strings"
	"time"

	"github.com/Gylmynnn/chat-cuy-be/internal/config"
	"github.com/Gylmynnn/chat-cuy-be/internal/model"
	"github.com/Gylmynnn/chat-cuy-be/internal/service"
	"github.com/Gylmynnn/chat-cuy-be/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userService service.UserService
	cfg         *config.Config
}

func NewUserHandler(userService service.UserService, cfg *config.Config) *UserHandler {
	return &UserHandler{
		userService: userService,
		cfg:         cfg,
	}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var req model.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if req.Username == "" || req.Email == "" || req.Password == "" {
		return utils.SendError(c, fiber.StatusBadRequest, "Username, email, and password are required")
	}

	resp, _, err := h.userService.Register(c.Context(), req)
	if err != nil {
		return utils.SendError(c, fiber.StatusConflict, err.Error())
	}

	token, err := utils.GenerateJWT(h.cfg.JWTSecret, resp.ID, resp.Username, resp.Email)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to generate token")
	}
	resp.Token = token

	return c.Status(fiber.StatusCreated).JSON(resp)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var req model.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if req.Email == "" || req.Password == "" {
		return utils.SendError(c, fiber.StatusBadRequest, "Email and password are required")
	}

	resp, err := h.userService.Login(c.Context(), req, h.cfg.JWTSecret)
	if err != nil {
		return utils.SendError(c, fiber.StatusUnauthorized, err.Error())
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

func (h *UserHandler) SearchUsers(c *fiber.Ctx) error {
	query := c.Query("username")
	if query == "" {
		return utils.SendError(c, fiber.StatusBadRequest, "Username query is required")
	}

	users, err := h.userService.SearchByUsername(c.Context(), query)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to search users")
	}

	return utils.SendSuccess(c, fiber.StatusOK, "Users found", users)
}

func (h *UserHandler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	user, err := h.userService.GetByID(c.Context(), userID)
	if err != nil {
		return utils.SendError(c, fiber.StatusNotFound, "User not found")
	}

	return utils.SendSuccess(c, fiber.StatusOK, "Profile retrieved", user)
}

func (h *UserHandler) UploadAvatar(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	file, err := c.FormFile("avatar")
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Avatar file is required")
	}

	// Validate file size (max 2MB)
	if file.Size > 2*1024*1024 {
		return utils.SendError(c, fiber.StatusBadRequest, "File size must be less than 2MB")
	}

	// Validate file type
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true}
	if !allowedExts[ext] {
		return utils.SendError(c, fiber.StatusBadRequest, "Only jpg, jpeg, png, gif, and webp files are allowed")
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d_%d%s", userID, time.Now().UnixNano(), ext)
	savePath := filepath.Join("uploads", "avatars", filename)

	// Save file
	if err := c.SaveFile(file, savePath); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to save file")
	}

	avatarURL := "/" + savePath

	// Update avatar in database
	if err := h.userService.UpdateAvatar(c.Context(), userID, avatarURL); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to update avatar")
	}

	// Return updated user
	user, err := h.userService.GetByID(c.Context(), userID)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to get user")
	}

	return utils.SendSuccess(c, fiber.StatusOK, "Avatar updated successfully", user)
}
