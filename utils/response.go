package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Success response: status, message, and data
func RespondSuccess(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": message,
		"data":    data,
	})
}

// Error response: status, message, and null data
func RespondError(c *gin.Context, message string, statusCode int) {
	c.JSON(statusCode, gin.H{
		"status":  "error",
		"message": message,
		"data":    nil,
	})
}
