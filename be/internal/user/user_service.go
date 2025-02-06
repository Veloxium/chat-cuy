package user

import (
	"context"
	"strconv"
	"time"

	"github.com/Gylmynnn/websocket-sesat/database"
	"github.com/Gylmynnn/websocket-sesat/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
)

const (
	secretKey = "sholatlimawaktu24434"
)

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

	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil

}



func (s *service) Login(c context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancle := context.WithTimeout(c, s.timeout)
	defer cancle()
	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserRes{}, err
	}

	err = utils.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserRes{}, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserRes{}, err
	}
	res := &LoginUserRes{AccessToken: ss, Username: u.Username, ID: strconv.Itoa(int(u.ID))}
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
		user = &User{
			Username: userRecord.DisplayName,
			Email:    userRecord.Email,
			Password: "",
		}

		user, err = s.Repository.CreateUser(ctx, user)
		if err != nil {
			return nil, err
		}
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, MyJWTClaims{
		ID:       strconv.Itoa(int(user.ID)),
		Username: user.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    strconv.Itoa(int(user.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	ss, err := jwtToken.SignedString([]byte(secretKey))
	if err != nil {
		return nil, err
	}

	res := &LoginUserWithGoogleRes{
		AccessToken: ss,
		ID:          strconv.Itoa(int(user.ID)),
		Username:    user.Username,
		Email:       user.Email,
	}

	return res, nil

}
