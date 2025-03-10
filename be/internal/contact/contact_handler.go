package contact

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


func (h *Handler) AddContact (c *gin.Context) {
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
