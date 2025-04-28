package chat

import (
	"context"
	"time"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateChat(ctx context.Context, req *CreateChatReq) (*CreateChatRes, error) {
	c, cancle := context.WithTimeout(ctx, s.timeout)
	defer cancle()
	chatReq := &Chat{
		ChatName: req.ChatName,
		IsGroup:  req.IsGroup,
	}
	res, err := s.Repository.CreateChat(c, chatReq)
	if err != nil {
		return nil, err
	}
	chatRes := &CreateChatRes{
		ChatName:  res.ChatName,
		IsGroup:   res.IsGroup,
		CreatedAt: res.CreatedAt,
	}
	return chatRes, nil
}

func (s *service) CreateMessage(ctx context.Context, message *CreateMessageReq) (*CreateMessageRes, error) {
	c, cancle := context.WithTimeout(ctx, s.timeout)
	defer cancle()

	// c := context.Background()

	msgReq := &Message{
		ChatID: message.ChatID,
		SenderID: message.SenderID,
		Content: message.Content,
	}

	res, err := s.Repository.CreateMessage(c, msgReq)
	if err != nil {
		return nil, err
	}
	msgRes := &CreateMessageRes{
		ID: res.ID,
		SenderID: res.SenderID,
		ChatID: res.ChatID,
		Content: res.Content,
		MessageType: res.MessageType,
		MediaURL: res.MediaURL,
		CreatedAt: res.CreatedAt,
	}
	return msgRes, nil
}

func (s *service) GetMessagesByChatID(ctx context.Context, chatID string) ([]*Message, error) {
	c, cancle := context.WithTimeout(ctx, s.timeout)
	defer cancle()
	res, err := s.Repository.GetMessagesByChatID(c, chatID)
	if err != nil {
		return nil, err
	}
	return res, nil
}
