package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/Gylmynnn/chat-cuy-be/internal/model"
	"github.com/Gylmynnn/chat-cuy-be/internal/repository"
	"github.com/Gylmynnn/chat-cuy-be/pkg/utils"
)

type UserService interface {
	Register(ctx context.Context, req model.RegisterRequest) (*model.LoginResponse, string, error)
	Login(ctx context.Context, req model.LoginRequest, jwtSecret string) (*model.LoginResponse, error)
	GetByID(ctx context.Context, id int) (*model.User, error)
	SearchByUsername(ctx context.Context, query string) ([]model.User, error)
	UpdateAvatar(ctx context.Context, userID int, avatarURL string) error
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo: userRepo}
}

func (s *userService) Register(ctx context.Context, req model.RegisterRequest) (*model.LoginResponse, string, error) {
	existing, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, "", err
	}
	if existing != nil {
		return nil, "", errors.New("email already registered")
	}

	existingUsername, err := s.userRepo.GetByUsername(ctx, req.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, "", err
	}
	if existingUsername != nil {
		return nil, "", errors.New("username already taken")
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, "", err
	}

	user := &model.User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, "", err
	}

	response := &model.LoginResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Bio:      user.Bio,
		Avatar:   user.Avatar,
	}

	return response, "", nil
}

func (s *userService) Login(ctx context.Context, req model.LoginRequest, jwtSecret string) (*model.LoginResponse, error) {
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := utils.GenerateJWT(jwtSecret, user.ID, user.Username, user.Email)
	if err != nil {
		return nil, err
	}

	return &model.LoginResponse{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Username,
		Bio:      user.Bio,
		Avatar:   user.Avatar,
		Token:    token,
	}, nil
}

func (s *userService) GetByID(ctx context.Context, id int) (*model.User, error) {
	return s.userRepo.GetByID(ctx, id)
}

func (s *userService) SearchByUsername(ctx context.Context, query string) ([]model.User, error) {
	return s.userRepo.SearchByUsername(ctx, query)
}

func (s *userService) UpdateAvatar(ctx context.Context, userID int, avatarURL string) error {
	return s.userRepo.UpdateAvatar(ctx, userID, avatarURL)
}
