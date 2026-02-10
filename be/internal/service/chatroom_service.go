package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Gylmynnn/chat-cuy-be/internal/model"
	"github.com/Gylmynnn/chat-cuy-be/internal/repository"
)

type ChatRoomService interface {
	GetUserChatRooms(ctx context.Context, userID int) ([]model.ChatRoomWithLastMessage, error)
	GetOrCreatePrivateRoom(ctx context.Context, userID1, userID2 int) (*model.ChatRoom, error)
}

type chatRoomService struct {
	chatRoomRepo repository.ChatRoomRepository
}

func NewChatRoomService(chatRoomRepo repository.ChatRoomRepository) ChatRoomService {
	return &chatRoomService{chatRoomRepo: chatRoomRepo}
}

func (s *chatRoomService) GetUserChatRooms(ctx context.Context, userID int) ([]model.ChatRoomWithLastMessage, error) {
	return s.chatRoomRepo.GetByUserID(ctx, userID)
}

func (s *chatRoomService) GetOrCreatePrivateRoom(ctx context.Context, userID1, userID2 int) (*model.ChatRoom, error) {
	room, err := s.chatRoomRepo.FindPrivateRoom(ctx, userID1, userID2)
	if err == nil {
		return room, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	newRoom := &model.ChatRoom{
		Name:    "",
		IsGroup: false,
	}
	if err := s.chatRoomRepo.Create(ctx, newRoom); err != nil {
		return nil, err
	}

	if err := s.chatRoomRepo.AddMember(ctx, newRoom.ID, userID1); err != nil {
		return nil, err
	}
	if err := s.chatRoomRepo.AddMember(ctx, newRoom.ID, userID2); err != nil {
		return nil, err
	}

	return newRoom, nil
}
