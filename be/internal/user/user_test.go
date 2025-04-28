package user

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/Gylmynnn/websocket-sesat/utils"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockUserService struct {
	mock.Mock
}

// FindUserByUserID implements Service.
func (m *MockUserService) FindUserByID(ctx context.Context, userID string) (*FindUserByIDRes, error) {
	args := m.Called(ctx, userID)
	if res, ok := args.Get(0).(*FindUserByIDRes); ok {
		return res, args.Error(0)
	}
	return nil, args.Error(1)
}

// SearchUsers implements Service.
func (m *MockUserService) SearchUsers(ctx context.Context, usernamePrefix string) ([]*SearchUsersRes, error) {
	args := m.Called(ctx, usernamePrefix)
	if res, ok := args.Get(0).([]*SearchUsersRes); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserService) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	args := m.Called(ctx, req)
	if res, ok := args.Get(0).(*CreateUserRes); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

// Login implements UserService.
func (m *MockUserService) Login(ctx context.Context, req *LoginUserReq) (*LoginUserRes, error) {
	args := m.Called(ctx, req)
	if res, ok := args.Get(0).(*LoginUserRes); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

// LoginWithFacebook implements UserService.
func (m *MockUserService) LoginWithFacebook(ctx context.Context, req *LoginUserWithFacebookReq) (*LoginUserWithFacebookRes, error) {
	args := m.Called(ctx, req)
	if res, ok := args.Get(0).(*LoginUserWithFacebookRes); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

// LoginWithGoogle implements UserService.
func (m *MockUserService) LoginWithGoogle(ctx context.Context, req *LoginUserWithGoogleReq) (*LoginUserWithGoogleRes, error) {
	args := m.Called(ctx, req)
	if res, ok := args.Get(0).(*LoginUserWithGoogleRes); ok {
		return res, args.Error(1)
	}
	return nil, args.Error(1)
}

func TestFindUserByID(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	handler := NewHundler(mockService)

	userIDParams := "6106c1afe1ad4af499724087"
	expectedRes := FindUserByIDRes{
		ID:             "6106c1afe1ad4af499724087",
		Username:       "anomaly",
		Email:          "anomaly24434@gmail.com",
		ProfilePicture: "test.png",
		AboutMessage:   "Hallo world",
	}

	// Set ekspektasi pemanggilan service
	mockService.On("FindUserByID", mock.Anything, userIDParams).Return(expectedRes, nil)

	// Simulasi request GET /api/user/:id
	req, _ := http.NewRequest(http.MethodGet, "/api/user/" + userIDParams, nil)

	fmt.Println("print find", req)
	w := httptest.NewRecorder()
	fmt.Println("print id httptest", w)

	// Setup context dan set param id
	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{
		{Key: "id", Value: userIDParams},
	}
	c.Request = req

	handler.FindUserByID(c)
	assert.Equal(t, http.StatusOK, w.Code)

	var res utils.ResFormatter
	err := json.Unmarshal(w.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, "find user successfully", res.Message)
	// Cek isi data
	data, _ := json.Marshal(res.Data)
	var user FindUserByIDRes
	json.Unmarshal(data, &user)

	assert.Equal(t, expectedRes.Username, user.Username)
	mockService.AssertExpectations(t)

}

func TestSearchUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	handler := NewHundler(mockService)

	usernameKey := "anomaly"
	expectedRes := []*SearchUsersRes{
		{
			ID:             "6106c1afe1ad4af499724087",
			Username:       "anomaly",
			Email:          "anomaly24434@gmail.com",
			ProfilePicture: "test.png",
			AboutMessage:   "Hallo world",
		},
	}

	mockService.On("SearchUsers", mock.Anything, usernameKey).Return(expectedRes, nil)

req, _ := http.NewRequest(http.MethodGet, "/api/user/search?username="+usernameKey, nil)
	fmt.Println("print search", req)
	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	// Jalankan handler
	handler.SearchUsers(c)
	// Assert
	assert.Equal(t, http.StatusOK, rr.Code)
	var res utils.ResFormatter
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, "search users successfully", res.Message)

	// Assert isi data
	jsonData, _ := json.Marshal(res.Data)
	var users []*SearchUsersRes
	err = json.Unmarshal(jsonData, &users)
	assert.NoError(t, err)
	assert.Len(t, users, 1)
	assert.Equal(t, expectedRes[0].Username, users[0].Username)

	mockService.AssertExpectations(t)

}

func TestCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	handler := NewHundler(mockService)

	userReq := CreateUserReq{
		Username: "anomaly",
		Email:    "anomaly24434@gmail.com",
		Password: "anomaly24434",
		ProfilePicture: "test.png",
		AboutMessage:   "Hallo world",
	}

	userRes := &CreateUserRes{
		ID:             "6106c1afe1ad4af499724087",
		Username:       "anomaly",
		Email:          "anomaly24434@gmail.com",
		ProfilePicture: "test.png",
		AboutMessage:   "Hallo world",
	}

	mockService.On("CreateUser", mock.Anything, &userReq).Return(userRes, nil)

	body, _ := json.Marshal(userReq)
	req, _ := http.NewRequest(http.MethodPost, "/api/register", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	handler.CreateUser(c)
	assert.Equal(t, http.StatusOK, rr.Code)

	var res utils.ResFormatter
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, "signup successfully", res.Message)
	mockService.AssertExpectations(t)
}

func TestLogin(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	handler := NewHundler(mockService)

	loginReq := LoginUserReq{
		Email:    "anomaly24434@gmail.com",
		Password: "anomaly24434",
	}

	loginRes := &LoginUserRes{
		AccessToken: "test=36106c1afe1ad4af499724087a3682de",
	}

	mockService.On("Login", mock.Anything, &loginReq).Return(loginRes, nil)

	body, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest(http.MethodPost, "/api/login/default", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	handler.Login(c)

	assert.Equal(t, http.StatusOK, rr.Code)
	var res utils.ResFormatter
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, "login successfully", res.Message)
	mockService.AssertExpectations(t)
}

func TestLoginWithFacebook(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	handler := NewHundler(mockService)

	loginReq := LoginUserWithFacebookReq{
		AccessToken: "fb=36106c1afe1ad4af499724087a3682de",
	}

	loginRes := &LoginUserWithFacebookRes{
		AccessToken: "test=36106c1afe1ad4af499724087a3682de",
	}

	mockService.On("LoginWithFacebook", mock.Anything, &loginReq).Return(loginRes, nil)

	body, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest(http.MethodPost, "/api/login/facebook", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	handler.LoginWithFacebook(c)

	assert.Equal(t, http.StatusOK, rr.Code)
	var res utils.ResFormatter
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, "login with facebook successfully", res.Message)
	mockService.AssertExpectations(t)
}

func TestLoginWithGoogle(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockService := new(MockUserService)
	handler := NewHundler(mockService)

	loginReq := LoginUserWithGoogleReq{
		AccessToken: "go=36106c1afe1ad4af499724087a3682de",
	}

	loginRes := &LoginUserWithGoogleRes{
		AccessToken: "test=36106c1afe1ad4af499724087a3682de",
	}

	mockService.On("LoginWithGoogle", mock.Anything, &loginReq).Return(loginRes, nil)

	body, _ := json.Marshal(loginReq)
	req, _ := http.NewRequest(http.MethodPost, "/api/login/google", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rr)
	c.Request = req

	handler.LoginWithGoogle(c)

	assert.Equal(t, http.StatusOK, rr.Code)
	var res utils.ResFormatter
	err := json.Unmarshal(rr.Body.Bytes(), &res)
	assert.NoError(t, err)
	assert.True(t, res.Success)
	assert.Equal(t, "login with google successfully", res.Message)
	mockService.AssertExpectations(t)
}
