package message

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func EntityException(c *gin.Context, message string) {
	res := Response{
		Status: http.StatusUnprocessableEntity,
		Message: message,
		Data: nil,
	}

	c.JSON(http.StatusUnprocessableEntity, res)
}