package message

import "github.com/gin-gonic/gin"

type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	Data interface{} `json:"data"`
}

func MessageSuccess(c *gin.Context, status int, message string, data interface{}) {
	res := Response{
		Status: status,
		Message: message,
		Data: data,
	}

	c.JSON(status, res)
}