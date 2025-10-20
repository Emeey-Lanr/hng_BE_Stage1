package handlers

import (
	"net/http"
	"stage1/models"
	"stage1/services"
	"stage1/utils"

	"github.com/gin-gonic/gin"
	// "github.com/go-playground/validator"
	// "log"
)

func AddAndAnalyseString(c *gin.Context){
	var req models.StringValue
 
	if err := c.ShouldBindJSON(&req); err != nil{
     services.RequestError(c, err)
	}

	hashedValue := services.CreateSHA256Hash(req.Value)

	_, exist :=  models.DB[hashedValue]

	if exist{
		utils.ErrorResponse(c, http.StatusConflict, "String already exist in the system")
		return
	}

	


}