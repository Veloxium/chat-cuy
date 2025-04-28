package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func BindFormAndValidate[T any](c *gin.Context) *T {
	var req T
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, ResFormatter{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "invalid form data:" + err.Error(),
			Data:       nil,
		})
		return nil
	}

	validate := validator.New()
	if err := validate.Struct(req); err != nil {
		errs := FormatValidationError(err)
		c.JSON(http.StatusBadRequest, ResFormatter{
			Success:    false,
			StatusCode: http.StatusBadRequest,
			Message:    "validation error",
			Data:       errs,
		})
		return nil
	}
	return &req
}
