package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func AppException(c *gin.Context, message string) {
	res := Response{
		Status: http.StatusInternalServerError,
		Message: message,
		Data: nil,
	}

	c.JSON(http.StatusInternalServerError, res)
}