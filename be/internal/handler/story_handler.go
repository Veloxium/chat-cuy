package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/Gylmynnn/chat-cuy-be/internal/model"
	"github.com/Gylmynnn/chat-cuy-be/internal/service"
	"github.com/Gylmynnn/chat-cuy-be/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

type StoryHandler struct {
	storyService service.StoryService
}

func NewStoryHandler(storyService service.StoryService) *StoryHandler {
	return &StoryHandler{storyService: storyService}
}

func (h *StoryHandler) CreateStory(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	file, err := c.FormFile("image")
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Image file is required")
	}

	// Validate file size (max 5MB)
	if file.Size > 5*1024*1024 {
		return utils.SendError(c, fiber.StatusBadRequest, "File size must be less than 5MB")
	}

	// Validate file type
	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := map[string]bool{".jpg": true, ".jpeg": true, ".png": true, ".gif": true, ".webp": true}
	if !allowedExts[ext] {
		return utils.SendError(c, fiber.StatusBadRequest, "Only jpg, jpeg, png, gif, and webp files are allowed")
	}

	// Generate unique filename
	filename := fmt.Sprintf("%d_%d%s", userID, time.Now().UnixNano(), ext)
	savePath := filepath.Join("uploads", "stories", filename)

	// Save file
	if err := c.SaveFile(file, savePath); err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to save file")
	}

	caption := c.FormValue("caption", "")

	story := &model.Story{
		UserID:   userID,
		ImageURL: "/" + savePath,
		Caption:  caption,
	}

	if err := h.storyService.Create(c.Context(), story); err != nil {
		// Clean up file on DB error
		os.Remove(savePath)
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to create story")
	}

	return utils.SendSuccess(c, fiber.StatusCreated, "Story created successfully", story)
}

func (h *StoryHandler) GetStories(c *fiber.Ctx) error {
	stories, err := h.storyService.GetAll(c.Context())
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to get stories")
	}

	if stories == nil {
		stories = []model.StoryWithUser{}
	}

	return c.Status(fiber.StatusOK).JSON(stories)
}

func (h *StoryHandler) GetMyStories(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	stories, err := h.storyService.GetByUserID(c.Context(), userID)
	if err != nil {
		return utils.SendError(c, fiber.StatusInternalServerError, "Failed to get stories")
	}

	if stories == nil {
		stories = []model.StoryWithUser{}
	}

	return c.Status(fiber.StatusOK).JSON(stories)
}

func (h *StoryHandler) DeleteStory(c *fiber.Ctx) error {
	userID := c.Locals("userID").(int)

	storyID, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return utils.SendError(c, fiber.StatusBadRequest, "Invalid story ID")
	}

	if err := h.storyService.Delete(c.Context(), storyID, userID); err != nil {
		return utils.SendError(c, fiber.StatusNotFound, "Story not found or unauthorized")
	}

	return utils.SendSuccess(c, fiber.StatusOK, "Story deleted successfully", nil)
}
