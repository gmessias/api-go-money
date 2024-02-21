package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MessageNotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, gin.H{
		"error": message,
	})
}

func MessageInternalError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": message,
	})
}

func MessageBadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"error": message,
	})
}

func MessageNoContentSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusNoContent, gin.H{
		"success": message,
	})
}
