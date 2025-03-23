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

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type UserEmail struct{
// 	Email  string `json:"email"`
// }
// type User struct{
//   Name string `json:"name"`
//   Age int`json:"age"`
//   Email *UserEmail `json:"email,omitempty"`
// }

// func main(){
// 	user1 := User{Name: "Prazo", Age: 34,Email: nil}
// 	user2:= User{Name: "Shabs", Age: 30,Email:&UserEmail{Email: "shabs76@gmail.com"}}

// 	jsonData1, err := json.Marshal(user1)
// 	if err != nil {
// 		log.Fatal("Failed to Marshal first user", err)
//		// no need to return because lo.fatal is automatic terminate program
// 	}
//   result1 := string(jsonData1)
//   fmt.Println("here is the result for the first data:", result1)

//   jsonData2,err := json.Marshal(user2)
//   if err != nil {
// 	log.Fatal("Failed to Marshal first user", err)
//
//   }
//   result2 := string(jsonData2)
//   fmt.Println("here is the result for the 2nd data:", result2)

// }

// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type Product struct{
// 	Name  string `json:"name"`
// 	Price float64 `json:"price"`
// }

// func main()  {
// 	product := `{"name":"Laptop", "price":34.5 }`

// 	var prd Product
// 	err := json.Unmarshal([]byte(product), (&prd))
// 	if err != nil {
// 		log.Fatal("Failed to Decode produjct", err)

// 	}
// 	fmt.Println(prd.Name, prd.Price,"$")

// }

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)



type Product struct{
	Name string `json:"name"`
	Price float64 `json:"price"`
}
type Response struct{
	Success  bool  `json:"success"`
	Message string `json:"message,omitempty"`
}
const Port = "3000"
func main()  {
	r :=mux.NewRouter()
	r.HandleFunc("/product",Productdetails).Methods("POST")	
	fmt.Println("server listen to port....",Port)
	log.Fatal(http.ListenAndServe(":"+Port, r))
}

func Productdetails(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	var prd Product
	err:= json.NewDecoder(r.Body).Decode(&prd)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	fmt.Println(prd.Name, prd.Price,"$")

	w.Header().Set("Content-Type", "Application/json")
	err = json.NewEncoder(w).Encode(Response{
		Message: "Successufly recive the data",
		Success: true,
	})
	if err != nil {
		log.Fatal("Failed to Encode data ", err)
	}
}

