package config

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// JSON serializes the api response properly to json
func JSON(c *gin.Context, message string, status int, data any) {
	responseData := gin.H{}

	switch data.(type) {
	case error:
		responseData = gin.H{
			"message": message,
			"errors":  data.(error).Error(),
			"status":  http.StatusText(status),
		}
	default:
		responseData = gin.H{
			"message": message,
			"data":    data,
			"status":  http.StatusText(status),
		}
	}
	c.JSON(status, responseData)
}

func HtmlRes(c *gin.Context, template string, status int, data gin.H) {
	c.HTML(status, template, data)
}
