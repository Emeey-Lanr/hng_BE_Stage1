package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"stage1/models"
	"stage1/utils"
	"strings"

	"github.com/gin-gonic/gin"
)


func RequestError (c *gin.Context, err error){


	var typeError *json.UnmarshalTypeError

	 switch  {
     // if type error
	 case errors.As(err, &typeError):
			utils.ErrorResponse(c, http.StatusUnprocessableEntity, `Invalid data type for "value" (must be string) `)
		return
	 default:
  	utils.ErrorResponse(c, http.StatusBadRequest, `Invalid request body or missing  or missing "value" field`)
		return
	 }
    


}


func IfStringExist (value string) (bool, models.Data){

  for _, val := range models.DB {
	if strings.ToLower(val.Value) == strings.ToLower(value) {
		return true, val
	}
  }
  return false, models.Data{}
}