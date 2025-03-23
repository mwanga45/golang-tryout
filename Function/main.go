package main

import (
	"encoding/json"
	"fmt"
    


)

func intSeq() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {

    nextInt := intSeq()

    fmt.Println(nextInt())
    fmt.Println(nextInt())
    fmt.Println(nextInt())

    newInts := intSeq()
    fmt.Println(newInts())
// How to handle the Json data in Golang 
    data := map[string]int{"apple":1, "mango":2,"Orange":3}

    fmt.Println(data)
// json.Marshal it convert the go format of data into the Go format 
    JsonData,_ := json.Marshal(data)
    fmt.Println(string(JsonData))

    for name,position:= range data{
        fmt.Printf("the name of fruits is %v and it position is  %v \n",name,position)
    }
// CONVERT BACK INTO GO FORMAT USING json.unmarshal
    JsonData2 := []byte(JsonData)

    var data2 map[string]int

    BacktoGo := json.Unmarshal(JsonData2, &data2)
    fmt.Println(BacktoGo)

    for name,position:= range data{
        fmt.Printf("thIS name of fruits is %v and it position is  %v \n",name,position)
    }

    // another example on how to convert the Json format into Go Format

    frontend := []byte(`{"apple":1, "mango":2,"Orange":3}`)
//   declare the slice which will later will be are pointer to hold the value byte of jsonData  
    var slice map[string]int

    // use json.unmarshing to return the Go json format  back from the Json data format 

    BacktoGo2 := json.Unmarshal(frontend,&slice)

   defer  fmt.Println(BacktoGo2)


    for name,position := range slice{
        fmt.Printf("the name of Fruit after the conversion is %v and it position after the conversion is %v \n", name, position)
    }

    data1,_:= json.Marshal(true)
    fmt.Print(string(data1)) 
    fmt.Println(slice)


}


// Go supports anonymous functions, which can form closures. Anonymous functions are useful when you want to define a function inline without having to name it.

// Run codeCopy code
// package main
// import "fmt"
// This function intSeq returns another function, which we define anonymously in the body of intSeq. The returned function closes over the variable i to form a closure.

// func intSeq() func() int {
//     i := 0
//     return func() int {
//         i++
//         return i
//     }
// }
// func main() {
// We call intSeq, assigning the result (a function) to nextInt. This function value captures its own i value, which will be updated each time we call nextInt.

//     nextInt := intSeq()
// See the effect of the closure by calling nextInt a few times.

//     fmt.Println(nextInt())
//     fmt.Println(nextInt())
//     fmt.Println(nextInt())
// To confirm that the state is unique to that particular function, create and test a new one.

//     newInts := intSeq()
//     fmt.Println(newInts())
// }
// $ go run closures.go
// 1
// 2
// 3
// 1
// The last feature of functions we’ll look at for now is recursion.

// Next example: Recursion.