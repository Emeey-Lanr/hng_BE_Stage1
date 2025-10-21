package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"stage1/models"
	"stage1/utils"
	"strings"
	"fmt"

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


func IfStringExist (value string) (bool, models.Data, int){

  for id, val := range models.DB {
	if strings.ToLower(val.Value) == strings.ToLower(value) {
		return true, val, id
	}
  }
  return false, models.Data{}, 0
}


 func FilterThroughQuery (query string) (map[string]interface{}, int, error){

	vowel := map[string]string{"first vowel":"a","second vowel":"e","third vowel":"i", "fourth vowel":"o", "fifth vowel":"u", "last vowel":"u",}

length := map[string]int{"1":1, "2":2, "3":3, "4":4, "5":5, "6":6, "7":7, "8":8, "9":9, "10":10, "11":11}
count := map[string]int {"single":1, "double":2, "triple":3}

letter := []string{"a","b","c","d","e","f","g","h","i","j", "k", "l", "m", "n", "o", "p","q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

queryNumb := 0
numb := &queryNumb

queryBox := make(map[string]interface{})
 
   
// query 1

	if strings.Contains(query, "word palindromic") {
	
queryBox["is_palindrome"] = true

		splitedQuery := strings.Split(query, " ") // splited to check if the second word is single, double, or tripple
        
		fmt.Println(splitedQuery)
		data, exist :=  count[splitedQuery[1]] // a map  that checks if it exist
		
		if !exist {
          return queryBox, 0, fmt.Errorf("Query Passed but resulted in conflicting filters")
		}
		queryBox["word_count"] = data
		*numb = 1
 
	}

	// query 2
	if strings.Contains(query, "strings longer than"){
			splitedQuery := strings.Split(query, " ") // split query to check if it has 1, 2, 3, 4, etc
			
			data, exist := length[splitedQuery[3]] 

			if !exist {
               return queryBox, 0, fmt.Errorf("Query Passed but resulted in conflicting filters")
			}

            queryBox["min_length"] = data
			*numb = 2
		
	}
	

	
// Query 3 Checks for vowel
	if strings.Contains(query, "palindromic strings that contain the"){
		queryBox["is_palindrome"] = true

		splitedQuery := strings.Split(query, " ") //split query to contain only one index of first vowel, second, etc
		   
	  data, exist := vowel[fmt.Sprintf("%s %s", splitedQuery[5], splitedQuery[6])] // if the val of the ending word is first, second, third vowel, etc

				   if !exist{
					 
					  if !exist {
               return queryBox, 0, fmt.Errorf("Query Passed but resulted in conflicting filters")
			}

				   }		   

			
				    queryBox["contains_character"] = data
					*numb = 3
				
		
	}



	//Query 4 checks if it contains  any ch
	if strings.Contains(query, "strings containing the letter"){
		splitedQuery := strings.Split(query, " ") // split string to check if it has a letter

		fmt.Println(splitedQuery[len(splitedQuery) - 1])

		for _, val := range letter{
			 if splitedQuery[len(splitedQuery) - 1] == val   {
                queryBox["contains_character"] = val
				*numb = 4
				break
			 }else{
				return queryBox,0,  fmt.Errorf("Query Passed but resulted in conflicting filters")
			 }	
		}
	  
	}


 return queryBox, queryNumb, nil	
	 
 }