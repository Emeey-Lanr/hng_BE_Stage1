package services



import (
		"encoding/json"
		"errors"
	"net/http"
	"stage1/utils"
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