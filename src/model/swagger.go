package model

type Swagger struct { 
	Swagger string 						`json:"swagger"` 
	Info 	map[string]interface{} 		`json:"info"` 
	Paths 	map[string]interface{} 		`json:"paths"` 
}