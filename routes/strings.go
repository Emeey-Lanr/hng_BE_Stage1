package routes

import (
	"stage1/handlers"

	"github.com/gin-gonic/gin"
)

func StringRoutes (r *gin.Engine) {
	r.POST("/strings", handlers.AddAndAnalyseString)
	r.GET("/strings/:string_value", handlers.GetSpecificString)
    r.GET("/strings")
 
    r.DELETE("/strings/:string_value", handlers.DeleteSpecificString)	
}