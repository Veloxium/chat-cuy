package utils

import (
	"strconv"
	"github.com/gin-gonic/gin"
)


func ParseIDParam(c *gin.Context, paramName string) (int64, error) {
	idStr := c.Param(paramName)
	return strconv.ParseInt(idStr, 10, 64)
}
