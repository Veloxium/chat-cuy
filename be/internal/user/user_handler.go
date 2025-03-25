package user

import (
	"github.com/Gylmynnn/websocket-sesat/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	Service
}

func NewHundler(s Service) *Handler {
	return &Handler{
		Service: s,
	}
}

func (h *Handler) CreateUser(c *gin.Context) {
	var u CreateUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.Service.CreateUser(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
		return
	}

	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "signup successfully",
		Data:       res,
	})
}

func (h *Handler) Login(c *gin.Context) {
	var u LoginUserReq
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.Service.Login(c.Request.Context(), &u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
		return
	}

	c.SetCookie("jwt", res.AccessToken, 3600, "/", "localhost", false, true)

	newResponse := &LoginUserRes{
		Username:       res.Username,
		ID:             res.ID,
		AccessToken:    res.AccessToken,
		Email:          res.Email,
		ProfilePicture: res.ProfilePicture,
		AboutMessage:   res.AboutMessage,
		CreatedAt:      res.CreatedAt,
	}

	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "login successfully",
		Data:       newResponse,
	})
}

func (h *Handler) LoginWithGoogle(c *gin.Context) {
	var user LoginUserWithGoogleReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.Service.LoginWithGoogle(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
		return
	}

	c.SetCookie("jwt", res.AccessToken, 3600, "/", "localhost", false, true)

	newResponse := &LoginUserWithGoogleRes{
		Email:          res.Email,
		Username:       res.Username,
		ID:             res.ID,
		AccessToken:    res.AccessToken,
		ProfilePicture: res.ProfilePicture,
		AboutMessage:   res.AboutMessage,
		CreatedAt:      res.CreatedAt,
	}

	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "login with google successfully",
		Data:       newResponse,
	})

}

func (h *Handler) LoginWithFacebook(c *gin.Context) {
	var user LoginUserWithFacebookReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.Service.LoginWithFacebook(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
		return
	}

	c.SetCookie("jwt", res.AccessToken, 3600, "/", "localhost", false, true)

	newResponse := &LoginUserWithFacebookRes{
		Email:          res.Email,
		Username:       res.Username,
		ID:             res.ID,
		AccessToken:    res.AccessToken,
		ProfilePicture: res.ProfilePicture,
		AboutMessage:   res.AboutMessage,
		CreatedAt:      res.CreatedAt,
	}

	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "login with facebook successfully",
		Data:       newResponse,
	})

}

func (h *Handler) Logout(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "logout successfully",
		Data:       nil,
	})
}
