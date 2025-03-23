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

// start convert Go struct into the Json 
user := User{Name: "Issa", Email: "prazoo45@gmail.com", Age:23}
if jsonData,err := json.Marshal(user);err != nil{
	fmt.Errorf("faild to marshail %v",err)
	return
}else{
	fmt.Println(string(jsonData))
}

}