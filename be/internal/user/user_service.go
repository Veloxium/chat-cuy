package user

import (
	"context"
	"database/sql"
	"strconv"
	"time"

	"github.com/Gylmynnn/websocket-sesat/database"
	"github.com/Gylmynnn/websocket-sesat/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey = utils.LoadENV("JWTSECRETKEY")

type MyJWTClaims struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

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

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancle := context.WithTimeout(c, s.timeout)
	defer cancle()

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	user := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
		Avatar:   "https://api.dicebear.com/9.x/adventurer/svg?seed=Aiden",
		Bio:      "Hallo world",
	}

	user, err = s.Repository.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:        strconv.FormatInt(user.ID, 10),
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Bio:       user.Bio,
		CreatedAt: user.CreatedAt,
	}

	return res, nil

}

func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancle := context.WithTimeout(c, s.timeout)
	defer cancle()
	user, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserRes{}, err
	}

	err = utils.CheckPassword(req.Password, user.Password)
	if err != nil {
		return &LoginUserRes{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.FormatInt(user.ID, 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserRes{}, err
	}
	res := &LoginUserRes{
		AccessToken: ss,
		Username:    user.Username,
		ID:          strconv.FormatInt(user.ID, 10),
		Email:       user.Email,
		Avatar:      user.Avatar,
		Bio:         user.Bio,
		CreatedAt:   user.CreatedAt,
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
				Username: userRecord.DisplayName,
				Email:    userRecord.Email,
				Password: "",
				Avatar:   userRecord.PhotoURL,
				Bio:      "Hallo world",
			}

			user, err = s.Repository.CreateUser(ctx, user)
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	}

	if user.ID == 0 {
		return nil, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.FormatInt(user.ID, 10),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.FormatInt(user.ID, 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	res := &LoginUserWithFacebookRes{
		AccessToken: ss,
		ID:          strconv.FormatInt(user.ID, 10),
		Username:    user.Username,
		Email:       user.Email,
		Avatar:      user.Avatar,
		Bio:         user.Bio,
		CreatedAt:   user.CreatedAt,
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
				Username: userRecord.DisplayName,
				Email:    userRecord.Email,
				Password: "",
				Avatar:   userRecord.PhotoURL,
				Bio:      "Hallo world",
			}

			user, err = s.Repository.CreateUser(ctx, user)
			if err != nil {
				return nil, err
			}
		}
	}

	if user.ID == 0 {
		return nil, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.FormatInt(user.ID, 10),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.FormatInt(user.ID, 10),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	res := &LoginUserWithGoogleRes{
		AccessToken: ss,
		ID:          strconv.FormatInt(user.ID, 10),
		Username:    user.Username,
		Email:       user.Email,
		Avatar:      user.Avatar,
		Bio:         user.Bio,
		CreatedAt:   user.CreatedAt,
	}

	return res, nil

}
