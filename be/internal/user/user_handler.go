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

func (h *Handler) UpdateUserByID(c *gin.Context) {
	userID := c.Param("id")

	req := utils.BindFormAndValidate[UpdateUserReq](c)
	if req == nil {
		return
	}

	fileHeader, err := c.FormFile("profile_picture")
	if err == nil {
		file, err := fileHeader.Open()
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ResFormatter{
				Success:    false,
				StatusCode: http.StatusInternalServerError,
				Message:    "failed to open file: " + err.Error(),
				Data:       nil,
			})
			return
		}
		defer file.Close()

		pictureURL, err := utils.UploadUserProfile(c.Request.Context(), file)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utils.ResFormatter{
				Success:    false,
				StatusCode: http.StatusInternalServerError,
				Message:    "upload to cloudinary error: " + err.Error(),
				Data:       nil,
			})
			return
		}
		req.ProfilePicture = &pictureURL
	}

	res, err := h.Service.UpdateUser(c.Request.Context(), userID, req)
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
		Message:    "update users successfully",
		Data:       res,
	})
}

func (h *Handler) FindUserByID(c *gin.Context) {
	userID := c.Param("id")
	res, err := h.Service.FindUserByID(c.Request.Context(), userID)
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
		Message:    "find user successfully",
		Data:       res,
	})
}

func (h *Handler) SearchUsers(c *gin.Context) {
	usernamePrefix := c.Query("username")
	res, err := h.Service.SearchUsers(c.Request.Context(), usernamePrefix)
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
		Message:    "search users successfully",
		Data:       res,
	})
}

func (h *Handler) CreateUser(c *gin.Context) {
	req := utils.BindFormAndValidate[CreateUserReq](c)
	if req == nil {
		return
	}

	res, err := h.Service.CreateUser(c.Request.Context(), req)
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
	req := utils.BindFormAndValidate[LoginUserReq](c)
	res, err := h.Service.Login(c.Request.Context(), req)
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

	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "login successfully",
		Data:       res,
	})
}

func (h *Handler) LoginWithGoogle(c *gin.Context) {
	req := utils.BindFormAndValidate[LoginUserWithGoogleReq](c)
	res, err := h.Service.LoginWithGoogle(c.Request.Context(), req)
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

	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "login with google successfully",
		Data:       res,
	})

}

func (h *Handler) LoginWithFacebook(c *gin.Context) {
	req := utils.BindFormAndValidate[LoginUserWithFacebookReq](c)
	res, err := h.Service.LoginWithFacebook(c.Request.Context(), req)
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

	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "login with facebook successfully",
		Data:       res,
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
