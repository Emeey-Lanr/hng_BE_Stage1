package models




type StringValue struct {
	Value string `json:"value" binding:"required"`
}


type Properties struct {
   Length int `json:"length"`
   Is_Palindrome  bool `json:"is_palindrome"`
   Unique_Characters int  `json:"unique_characters"`
   Word_Count int `json:"word_count"`
   Sha256_Hash string `json:"sha256_hash"`
   Character_Frequency map[string]int `json:"character_frequency"`
}

type Data struct {
	Id string `json:"id"`
	Value string `json:"value"`
	Properties Properties `json:"properties"`
	Created_at string `json:"created_at"`

}



var DB  []Data


type StringFiltering struct {
	Is_Palindrome *bool `form:"is_palindrome"`
	Min_Length *int `form:"min_length"`
	Max_Length *int `form:"max_length"`
	Word_Count *int `form:"word_count"`
	Contains_Character *string `form:"contains_character"`

}


type GetStringsThroughQueryResponse struct{
	Data []Data `json:"data"`
	Count int `json:"count"`
	Filters_Applied StringFiltering `json:"filters_applied"`
}