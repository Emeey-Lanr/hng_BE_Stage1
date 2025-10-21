package services

import (
	"crypto/sha256"
	"fmt"
	"stage1/models"

	"strings"
	"unicode"
)

func CreateSHA256Hash (value string) string {
    hash := sha256.Sum256([]byte(value))

	
	return fmt.Sprintf("%x", hash)
}

// is palindrome 


func CleanString (value string) string {
    var cleaned strings.Builder //+=

	for _, r := range value{ // iterate over  runes
    if unicode.IsLetter(r){  // keeps only letter
      cleaned.WriteRune(unicode.ToUpper(r))
	}
	}

	return cleaned.String()
}




func CheckIfPalindrome (value string) bool {
  
   var stringFromTheBack string

  for i := 0; i < len(value); i ++{
	// reverse string
		
	 stringFromTheBack +=   strings.Split(value, "")[len(value) - ( i + 1)]
		
	
 }

  fmt.Println(stringFromTheBack, value)
if stringFromTheBack != value{
	return false
}


return true

}


func UniqueCharacter (value string) (map[string]int, int) {
	
	var uniCha = make(map[string]int)

	 for _, val := range value {
		// converted rune val into string
	     convertedRuneToString := string(val)
          
		 //changed case to lower
		 lowerCaseValue := strings.ToLower(convertedRuneToString)

		uniCha[lowerCaseValue]++
	 }

   return uniCha, len(uniCha)
 }




 func Filter (query models.StringFiltering) []models.Data{
	
    var result []models.Data

	for _, val := range models.DB {
       palindrome := val.Properties.Is_Palindrome == *query.Is_Palindrome
	   length := val.Properties.Length >= *query.Min_Length && val.Properties.Length <= *query.Max_Length
	   word_count := val.Properties.Word_Count == *query.Word_Count
	   _, contains_characters := val.Properties.Character_Frequency[strings.ToLower(*query.Contains_Character)]

	   if palindrome && length && word_count && contains_characters {
         result = append(result, val)
	   }
	}

	return result
 }