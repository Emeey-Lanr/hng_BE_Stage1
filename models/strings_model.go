package models




type StringValue struct {
	Value string `json:"value" binding:"required"`
}



type Data struct {
	Id string `json:"id"`
	Value string `json:"value"`
	Properties interface{} `json:"properties"`
	Created_at string `json:"created_at"`

}

var DB  map[string]Data