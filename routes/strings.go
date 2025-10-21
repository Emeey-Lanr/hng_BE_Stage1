package routes

import (
	"stage1/handlers"

	"github.com/gin-gonic/gin"
)

func StringRoutes (r *gin.Engine) {
	r.POST("/strings", handlers.AddAndAnalyseString)
	r.GET("/strings/:string_value", handlers.GetSpecificString)
    r.GET("/strings", handlers.GetStringsThroughQuery)
    r.GET("/strings/filter-by-natural-language", handlers.FilterThroughNaturalLanguage)
    r.DELETE("/strings/:string_value", handlers.DeleteSpecificString)	
}