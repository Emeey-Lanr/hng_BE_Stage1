package handlers

import (
	"net/http"
	"stage1/models"
	"stage1/services"
	"stage1/utils"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func AddAndAnalyseString(c *gin.Context){
	var req models.StringValue

 
	if err := c.ShouldBindJSON(&req); err != nil{
     services.RequestError(c, err)
	}

	hashedValue := services.CreateSHA256Hash(req.Value)
    
	// chest if string exists in the in memory db
	 exist, _ := services.IfStringExist(req.Value)

	 if exist{
         utils.ErrorResponse(c, http.StatusConflict, "String already exists in the system")
	 }

	
	// length of string
	characters := []rune(req.Value)

	//string without spaces
	cleanedString := services.CleanString(req.Value)

	// returns true or false
	is_Palindrome := services.CheckIfPalindrome(cleanedString)
     
	// return the frequency counts and the length of the unique characters
	frequency, lengthOfSpecialCharacters := services.UniqueCharacter(req.Value)

	//Slice of words sperated by space
	wordCount := strings.Fields(req.Value)

	
	properties := models.Properties{Length: len(characters),
		 Is_Palindrome: is_Palindrome, Unique_Characters:lengthOfSpecialCharacters, Word_Count: len(wordCount), Sha256_Hash: hashedValue, Character_Frequency: frequency}

     
    dataCreated := models.Data{Id: hashedValue, Value: req.Value, 
	Properties: properties, Created_at: time.Now().UTC().Format(time.RFC3339)}		 

	
   models.DB = append(models.DB, dataCreated)

  utils.SuccessResponse(c, http.StatusCreated, dataCreated)	

}


func GetSpecificString (c * gin.Context) {
  stringValue := c.Param("value")
 
  exist, value := services.IfStringExist(stringValue)
  
  if !exist {
	utils.ErrorResponse(c, http.StatusNotFound, "String does not exist in the system")
	return
  }

  utils.SuccessResponse(c, http.StatusOK, value)

}