package service

import (
	"context"

	"github.com/Gylmynnn/chat-cuy-be/internal/model"
	"github.com/Gylmynnn/chat-cuy-be/internal/repository"
)

type MessageService interface {
	SendMessage(ctx context.Context, senderID, chatRoomID int, text string) (*model.Message, error)
	GetMessages(ctx context.Context, chatRoomID int) ([]model.Message, error)
}

type messageService struct {
	messageRepo repository.MessageRepository
}

func NewMessageService(messageRepo repository.MessageRepository) MessageService {
	return &messageService{messageRepo: messageRepo}
}

func (s *messageService) SendMessage(ctx context.Context, senderID, chatRoomID int, text string) (*model.Message, error) {
	msg := &model.Message{
		Text:       text,
		SenderID:   senderID,
		ChatRoomID: chatRoomID,
	}

	if err := s.messageRepo.Create(ctx, msg); err != nil {
		return nil, err
	}

	return msg, nil
}

func (s *messageService) GetMessages(ctx context.Context, chatRoomID int) ([]model.Message, error) {
	return s.messageRepo.GetByChatRoomID(ctx, chatRoomID)
}
