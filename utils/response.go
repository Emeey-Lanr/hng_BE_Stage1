package utils

import (
	"github.com/gin-gonic/gin"
)
func ErrorResponse (c *gin.Context, method int, message string) {
   c.JSON(method, gin.H{"error":message})
}

func SuccessResponse (c *gin.Context, method int, data interface{}){
	c.JSON(method, data)
} 