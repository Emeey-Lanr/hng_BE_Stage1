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
   Character_Frequency interface {} `json:"character_frequency"`
}

type Data struct {
	Id string `json:"id"`
	Value string `json:"value"`
	Properties interface{} `json:"properties"`
	Created_at string `json:"created_at"`

}



var DB  []Data