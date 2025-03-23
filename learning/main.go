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

import (
	"encoding/json"
	"fmt"
	"log"
)
 

type UserEmail struct{
	Email  string `json:"email"`
}
type User struct{
  Name string `json:"name"`
  Age int`json:"age"`
  Email *UserEmail `json:"email,omitempty"`
}

func main(){
	user1 := User{Name: "Prazo", Age: 34,Email: nil}
	user2:= User{Name: "Shabs", Age: 30,Email:&UserEmail{Email: "shabs76@gmail.com"}}

	jsonData1, err := json.Marshal(user1)
	if err != nil {
		log.Fatal("Failed to Marshal first user", err)
		return
	}
  result1 := string(jsonData1)
  fmt.Println("here is the result for the first data:", result1)

  jsonData2,err := json.Marshal(user2)
  if err != nil {
	log.Fatal("Failed to Marshal first user", err)
		return
  }
  result2 := string(jsonData2)
  fmt.Println("here is the result for the 2nd data:", result2)
  

}