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




 func Filter(query models.StringFiltering) []models.Data {
	var result []models.Data

	for _, val := range models.DB {
		// Skip if Is_Palindrome filter is provided and doesn't match
		if query.Is_Palindrome != nil && val.Properties.Is_Palindrome != *query.Is_Palindrome {
			continue
		}

		// Skip if Min_Length is provided and string is shorter
		if query.Min_Length != nil && val.Properties.Length < *query.Min_Length {
			continue
		}

		// Skip if Max_Length is provided and string is longer
		if query.Max_Length != nil && val.Properties.Length > *query.Max_Length {
			continue
		}

		// Skip if Word_Count filter is provided and doesn't match
		if query.Word_Count != nil && val.Properties.Word_Count != *query.Word_Count {
			continue
		}

		// Skip if Contains_Character filter is provided and not found in string
		if query.Contains_Character != nil {
			char := strings.ToLower(*query.Contains_Character)
			if _, exists := val.Properties.Character_Frequency[char]; !exists {
				continue
			}
		}

		// Passed all applicable filters → add to result
		result = append(result, val)
	}

	return result
}


func FilterThroughNaturalLanguage (filter map[string]interface{}, numb int)([]models.Data){

  var result []models.Data



  for _, val := range models.DB {


	if numb == 1{
		wordCount :=  filter["word_count"].(int)
		if val.Properties.Is_Palindrome && val.Properties.Word_Count == int(wordCount) {
			result = append(result, val)
			return result
		}
	} 

	if numb == 2 {
		length := filter["min_length"].(int)
       if val.Properties.Length >= int(length){
		result = append(result, val)
		return result
	   }
	}

	if numb == 3 || numb == 4 {
        letter :=  filter["contains_character"].(string)
		_, exist := val.Properties.Character_Frequency[letter]

		if val.Properties.Is_Palindrome && exist {
			result = append(result, val)
			return result
		}
	}

	
  }

  return result

}
