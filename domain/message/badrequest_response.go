package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func BadRequestException(c *gin.Context, message string) {
	res := Response{
		Status: http.StatusBadRequest,
		Message: message,
		Data: nil,
	}

	c.JSON(http.StatusBadRequest, res)
}