package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
)

func main() {
	// concept map function we will use make to create map

	language := make(map[string]string)

	language["python"] = "Simple languaage"
	language["Javascript"] = "I love it  and it good"
	language["Golang"] = "i want to know much  i have that interest"

	fmt.Println(language)

	// using  for loop to interpret this map 

	for key,value := range language{
		fmt.Printf("How u fill use %v for me  %v \n", key, value)
	}
	/*
	if  you only want to display value 
	for _,value := range language{
		fmt.Printf("How u fill use %v for me  %v \n", key, value)
	}
    */

 token := generateToken(5)
 fmt.Println(token)
	
}
// create are simple token 

func generateToken(length int)string  {

	bytes := make([]byte,length);

	if _,err := rand.Read(bytes); err != nil{
		log.Fatal("Failed to generate Token ", err)
	}
	 return base64.URLEncoding.EncodeToString(bytes)
}