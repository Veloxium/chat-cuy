package user

import (
	"net/http"

	"github.com/Gylmynnn/websocket-sesat/pkg/utils"
	"github.com/gin-gonic/gin"
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

	c.SetCookie("jwt", res.accessToken, 3600, "/", "localhost", false, true)

	newRes := &LoginUserRes{
		Username: res.Username,
		ID:       res.ID,
	}

	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "login successfully",
		Data:       newRes,
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
