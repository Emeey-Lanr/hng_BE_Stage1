package handlers

import (
	"fmt"
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
	 exist, _, _ := services.IfStringExist(req.Value)

	 if exist{
         utils.ErrorResponse(c, http.StatusConflict, "String already exists in the system")
		 return
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
  stringValue := c.Param("string_value")
 
  exist, value, _ := services.IfStringExist(stringValue)
  
  if !exist {
	utils.ErrorResponse(c, http.StatusNotFound, "String does not exist in the system")
	return
  }

  utils.SuccessResponse(c, http.StatusOK, value)

}


func GetStringsThroughQuery(c *gin.Context) {
    var queryData models.StringFiltering

    // Bind query parameters
    if err := c.ShouldBindQuery(&queryData); err != nil {
        utils.ErrorResponse(c, http.StatusBadRequest, "Invalid query parameter values or types")
        return
    }

    // Filter data
    dataFound := services.Filter(queryData)

    // Build filters_applied map dynamically
    filtersApplied := make(map[string]interface{})
    if queryData.Is_Palindrome != nil {
        filtersApplied["is_palindrome"] = *queryData.Is_Palindrome
    }
    if queryData.Min_Length != nil {
        filtersApplied["min_length"] = *queryData.Min_Length
    }
    if queryData.Max_Length != nil {
        filtersApplied["max_length"] = *queryData.Max_Length
    }
    if queryData.Word_Count != nil {
        filtersApplied["word_count"] = *queryData.Word_Count
    }
    if queryData.Contains_Character != nil && *queryData.Contains_Character != "" {
        filtersApplied["contains_character"] = *queryData.Contains_Character
    }

    // Prepare response
    response := gin.H{
        "data":            dataFound,
        "count":           len(dataFound),
        "filters_applied": filtersApplied,
    }

    utils.SuccessResponse(c, http.StatusOK, response)
}


func FilterThroughNaturalLanguage (c *gin.Context){
  query := c.Query("query")

 if query == ""{
		utils.ErrorResponse(c, http.StatusBadRequest, "Unable to parse natural query")
		return
 }

 fmt.Println(query)

  filter, numb, err := services.FilterThroughQuery(strings.ToLower(query)) // the query I will use for the search
  
  if err != nil {
	utils.ErrorResponse(c, http.StatusUnprocessableEntity, err.Error())
	return
  }
  

  if len(filter) < 1 {
	utils.ErrorResponse(c, http.StatusBadRequest, "Unable to parse natural query")
	return
  }


  filterResult :=  services.FilterThroughNaturalLanguage(filter, numb) // the search result

  if len(filterResult) < 1 {
      utils.SuccessResponse(c, http.StatusNoContent, "")
  }
  
  data := models.NaturalLanguageResponse{Data: filterResult,
	 Count: len(filterResult), Interpreted_Query: models.InterpretedQuery{Original:query, Parsed_Filter: filter } }

  utils.SuccessResponse(c, http.StatusOK, data)
  

}



func DeleteSpecificString (c *gin.Context){
	stringValue := c.Param("string_value")

	exist, _, id := services.IfStringExist(stringValue)

	if !exist {
	utils.ErrorResponse(c, http.StatusNotFound, "String does not exist in the system")
	return
  }
	 
	models.DB = append(models.DB[:id], models.DB[id+1:]... )

	
  utils.SuccessResponse(c, http.StatusNoContent, models.Data{})

}
