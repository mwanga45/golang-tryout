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
JsonTstruct()
MapTjson()
JsonTMap()

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

// start covert json to go struct
func JsonTstruct()(string, error){
	user := `{"name":"issa","email":"prazoo45@gmail.com","age":23}`
	// user := `{"name":"Issa","email":"prazoo45@gmail.com","age":23}` 
	var details User
	err := json.Unmarshal([]byte(user),&details)
	if err != nil{
		return "", fmt.Errorf("failed to unmarshal json", err)	
	}
	result := fmt.Sprintf("Name: %s, Email: %s, Age: %d", details.Name, details.Email, details.Age)
	fmt.Println(result)
   return result, nil
}
// convert map go into json

func MapTjson()(string, error){
	data := map[string]interface{}{
		"Age":23,
		"Email":"Prazzo45@gmail.com",
		"Name":"Issa",

	}
	jsonData, err := json.Marshal(data)
	if err != nil{
		return "",fmt.Errorf("Failed to Marshal Map",err)
	}
	result := string(jsonData)
	fmt.Println(result)
	return result, nil
}
// convert json to Map go

func JsonTMap()(string, error){
var data map[string] interface{}
user := `{"name":"Issa", "email":"Prazoo45@gmail.com", "age":23}`
err := json.Unmarshal([]byte(user), &data)
if err != nil{
	return "", fmt.Errorf("Failed to Decode the user",err)
}
result := fmt.Sprintf("name is %s email is %s and age is %d",data["name"], data["email"], data["age"])
fmt.Println(result)
return result, nil
}