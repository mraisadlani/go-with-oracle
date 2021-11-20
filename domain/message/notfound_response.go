package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NotFoundException(c *gin.Context, message string) {
	res := Response{
		Status: http.StatusNotFound,
		Message: message,
		Data: nil,
	}

	c.JSON(http.StatusNotFound, res)
}