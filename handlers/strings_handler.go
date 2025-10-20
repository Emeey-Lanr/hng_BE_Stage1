package handlers

import (
	"errors"
	"net/http"
	"stage1/models"
	"stage1/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func AddAndAnalyseString(c *gin.Context){
	var stringValue models.StringValue
 
	if err := c.ShouldBindJSON(&stringValue); err != nil{

     var validate validator.ValidationErrors

	//  bad request body missing value
	if errors.As(err, &validate){
      utils.ErrorResponse(c, http.StatusBadRequest, `Invalid request body or missing "value" field`)
	 return
	}

	// if value is not string
	utils.ErrorResponse(c, http.StatusUnprocessableEntity, `Invalid data type for "value" (must be string)`)
	return

	}



}