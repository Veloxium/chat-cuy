package contact

import (
	"net/http"

	"github.com/Gylmynnn/websocket-sesat/utils"
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

func (h *Handler) DeleteContact(c *gin.Context) {
	contactID, err := utils.ParseIDParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
		c.Abort()
		return
	}

	err = h.Service.DeleteContact(c.Request.Context(), contactID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "contact deleted successfully",
		Data:       nil,
	})
}

func (h *Handler) GetAllContacts(c *gin.Context) {
	res, err := h.Service.GetAllContacts(c.Request.Context())
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
		Message:    "get contacts successfully",
		Data:       res,
	})
}

func (h *Handler) GetContactByID(c *gin.Context) {
	contactID, err := utils.ParseIDParam(c, "id")
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
		c.Abort()
		return
	}
	res, err := h.Service.GetContactByID(c.Request.Context(), contactID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusInternalServerError,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, utils.ResFormatter{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    "get contact by id add successfully",
		Data:       res,
	})
}

func (h *Handler) AddContact(c *gin.Context) {
	var contact CreateContactReq
	if err := c.ShouldBindJSON(&contact); err != nil {
		c.JSON(http.StatusBadRequest, utils.ResFormatter{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "error :" + err.Error(),
			Data:       nil,
		})
	}

	res, err := h.Service.AddContact(c.Request.Context(), &contact)
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
		Message:    "contact add successfully",
		Data:       res,
	})

}
