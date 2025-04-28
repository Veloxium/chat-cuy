package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/Gylmynnn/websocket-sesat/database"
	"github.com/Gylmynnn/websocket-sesat/utils"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = utils.LoadENV("JWTSECRETKEY")

type service struct {
	Repository
	timeout time.Duration
}

// UpdateUser implements Service.

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) UpdateUser(ctx context.Context, userID string, req *UpdateUserReq) (*UpdateUserRes, error) {
	c, cancle := context.WithTimeout(ctx, s.timeout)
	defer cancle()

	userReq := &User{}
	if req.Username != nil {
		userReq.Username = *req.Username
	}
	if req.ProfilePicture != nil {
		userReq.ProfilePicture = *req.ProfilePicture
	}
	if req.AboutMessage != nil {
		userReq.AboutMessage = *req.AboutMessage
	}

	res, err := s.Repository.Putser(c, userID, userReq)
	if err != nil {
		return nil, err
	}

	userRes := &UpdateUserRes{
		ID:             res.ID,
		UpdatedAt:      res.UpdatedAt,
	}

	return userRes, nil
}

func (s *service) FindUserByID(ctx context.Context, userID string) (*FindUserByIDRes, error) {
	c, cancle := context.WithTimeout(ctx, s.timeout)
	defer cancle()

	user, err := s.Repository.GetUserByID(c, userID)
	if err != nil {
		return nil, err
	}

	userRes := &FindUserByIDRes{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
		AboutMessage:   user.AboutMessage,
	}

	return userRes, nil
}

func (s *service) SearchUsers(ctx context.Context, usernamePrefix string) ([]*SearchUsersRes, error) {
	// if len(usernamePrefix) == 0 || usernamePrefix[0] != '@' {
	// 	return nil, errors.New("username must start with '@'")
	// }

	c, cancle := context.WithTimeout(ctx, s.timeout)
	defer cancle()

	users, err := s.Repository.SearchUsersByUsername(c, usernamePrefix)
	if err != nil {
		return nil, err
	}

	var result []*SearchUsersRes
	for _, user := range users {
		result = append(result, &SearchUsersRes{
			ID:             user.ID,
			Username:       user.Username,
			Email:          user.Email,
			ProfilePicture: user.ProfilePicture,
			AboutMessage:   user.AboutMessage,
		})
	}

	return result, nil
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancle := context.WithTimeout(c, s.timeout)
	defer cancle()

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username:       req.Username,
		Email:          req.Email,
		Password:       hashedPassword,
		ProfilePicture: "https://api.dicebear.com/9.x/adventurer/svg?seed=Aiden",
		AboutMessage:   "hi there i am using chat cuyyy",
	}

	user, err = s.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
		AboutMessage:   user.AboutMessage,
		CreatedAt:      user.CreatedAt,
	}

	return res, nil

}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancle := context.WithTimeout(c, s.timeout)
	defer cancle()
	user, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		return nil, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, utils.MyJWTClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}
	res := &LoginUserRes{
		AccessToken:    ss,
		Username:       user.Username,
		ID:             user.ID,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
		AboutMessage:   user.AboutMessage,
		CreatedAt:      user.CreatedAt,
	}
	return res, nil
}

func (s *service) LoginWithFacebook(c context.Context, req *LoginUserWithFacebookReq) (*LoginUserWithFacebookRes, error) {
	ctx, cancle := context.WithTimeout(c, s.timeout)
	defer cancle()

	token, err := database.AuthClient.VerifyIDToken(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	userRecord, err := database.AuthClient.GetUser(ctx, token.UID)
	if err != nil {
		return nil, err
	}

	user, err := s.Repository.GetUserByEmail(ctx, userRecord.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			user = &User{
				Username:       userRecord.DisplayName,
				Email:          userRecord.Email,
				Password:       "",
				ProfilePicture: userRecord.PhotoURL,
				AboutMessage:   "Hallo cuyyy cihuyyy",
			}

			user, err = s.Repository.CreateUser(ctx, user)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	if len(user.ID) == 0 {
		return nil, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, utils.MyJWTClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	res := &LoginUserWithFacebookRes{
		AccessToken:    ss,
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
		AboutMessage:   user.AboutMessage,
		CreatedAt:      user.CreatedAt,
	}
	return res, nil

}

func (s *service) LoginWithGoogle(c context.Context, req *LoginUserWithGoogleReq) (*LoginUserWithGoogleRes, error) {

	ctx, cancle := context.WithTimeout(c, s.timeout)
	defer cancle()

	token, err := database.AuthClient.VerifyIDToken(ctx, req.AccessToken)
	if err != nil {
		return nil, err
	}

	userRecord, err := database.AuthClient.GetUser(ctx, token.UID)
	if err != nil {
		return nil, err
	}

	user, err := s.Repository.GetUserByEmail(ctx, userRecord.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			user = &User{
				Username:       userRecord.DisplayName,
				Email:          userRecord.Email,
				Password:       "",
				ProfilePicture: userRecord.PhotoURL,
				AboutMessage:   "Hallo cuyyy cihuyyy",
			}

			user, err = s.Repository.CreateUser(ctx, user)
			if err != nil {
				return nil, err
			}
		}
	}

	if len(user.ID) == 0 {
		return nil, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, utils.MyJWTClaims{
		ID:       user.ID,
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    user.ID,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	res := &LoginUserWithGoogleRes{
		AccessToken:    ss,
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		ProfilePicture: user.ProfilePicture,
		AboutMessage:   user.AboutMessage,
		CreatedAt:      user.CreatedAt,
	}

	return res, nil

}
