// package main

// import (
// 	"encoding/json"
// 	"fmt"
// )

// type User struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

// func main() {
// 	user := User{Name: "Issa", Age: 24}
// 	jsonData,err:= json.Marshal(user)
// 	if err != nil{
// 		fmt.Println("Somthing went wrong", err)
// 		return
// 	}
// 	fmt.Println(string(jsonData))
// }

package main
 

type UserEmail struct{
	Email  string `json:"email"`
}
type User struct{
  Name string `json:"name"`
  Age string `json:"age"`
  Email *UserEmail `json:"email"`
}

func main(){
	
}