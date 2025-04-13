package main

import (
	"fmt"
	"reflect"
	"sort"

)

func main() {
	FirstArray := []string{"Football", "Baskatball", "mango"}
	fmt.Println("Orignal arrey",FirstArray)
	//to add the value or item within arrey
	FirstArray = append(FirstArray, "watermelo", "lemon", "Orange")
	fmt.Println("After append arrey",FirstArray)
   // to remove it from arreys the  count is not account last value in onsideration  
   // so this example it only remove  value at range of 0 - 1  but last value is not put in account 
   FirstArray = append(FirstArray[1:])
   fmt.Println("After remove arrey",FirstArray)
   // here it will start at range of  1 up to 3 but the 3 index it value it will remain as it is 
   FirstArray = append(FirstArray[1:3])
   fmt.Println("After remove arrey",FirstArray)

   // use make() for memory  management 

   SecondArray := make([]int, 4)
   SecondArray[0]= 234
   SecondArray[1]= 134
   SecondArray[2]= 334
   SecondArray[3]= 434

   SecondArray = append(SecondArray, 23, 34,53,21)
    fmt.Println(SecondArray)
	// here try to arrange the value from the highest to lowest 
	SecondArray = sort.IntSlice(SecondArray)
	fmt.Println(SecondArray)
	fmt.Println(len(SecondArray))



    collection1 := []int{1,3,2,4}
    collection2 := []int{3,1,4,2}
    equal1 := reflect.DeepEqual(collection1,collection2)
    fmt.Println(equal1)

    sort.Ints(collection1)
    sort.Ints(collection2)
    fmt.Println("1st collection", collection1)
    fmt.Println("2nd collection", collection2)
    
    equal := reflect.DeepEqual(collection1,collection2)
    fmt.Println(equal)
    collection2 = append(collection2, 5 , 6, 7 )
    fmt.Println(collection1)
    fmt.Println("2nd collection", collection2)
    index := 3
    // collection1 = append(collection1[:index], collection1[index + 1])
    // fmt.Println(collection1)
    collection3 := []int{10, 11, 12, 13, 14, 16, 15}
    sort.Ints(collection3)
    collection3 = append(collection3[:index] , collection3[index + 1:]...)
    fmt.Println(collection3)
   

   var str = "Hello world"
    bytes := []byte(str)
   
   fmt.Println("bytes slice: ",bytes)
}
// Username := "Issa mwanga"
// Password := "1234567" 
var SecreteKey = []byte("secrete-key")
// func createToken(Username string, Password string)(string, error){
//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.NewWithClaims{
//         "Username": Username,
//         "Password": Password,
//         "exp": time.Now().Add(time.Hour * 24).Unix(),
//     })
//       token, err :=  token.SignedString(&SecreteKey) 
//       if  err != nil{
//           fmt.Println("failed to load  token in slice")
//         return err
//      }
//     return token 
// }