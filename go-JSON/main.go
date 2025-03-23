package main

import (
	"encoding/json"
	"fmt"

)

// create user struct
type User struct{
	Name string `json:"name"`
	Email string `json:"email"`
	Age int `json:"age"`
}
func main() {
fmt.Println("learing golang/json")
StructTJson()

}
// start convert Go struct into the Json 
func StructTJson()(string, error){
  user := User{Name: "Issa", Email: "Prazoo45@gmail.com",Age: 23}
  jsonData, err := json.Marshal(user)
  if err != nil{
	return "",fmt.Errorf("Failed to Marshal struct %v",err)
  }
result := string(jsonData)
fmt.Println(result)
return result, nil
	
}