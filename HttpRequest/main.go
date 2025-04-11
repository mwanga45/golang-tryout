package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)
type ExpoPushMessage struct {
	To    string `json:"to"`
	Title string `json:"title"`
	Body  string `json:"body"`
	Sound string `json:"sound,omitempty"`
}
// this we try to get request from the web browser using package net/HTTP
const url = "https://www.google.com/search?q=nida+online&oq=nida&gs_lcrp=EgZjaHJvbWUqDAgBECMYJxiABBiKBTIMCAAQRRg5GLEDGIAEMgwIARAjGCcYgAQYigUyBwgCEAAYgAQyBwgDEAAYgAQyBwgEEAAYgAQyBwgFEAAYgAQyDQgGEC4YrwEYxwEYgAQyBwgHEAAYgAQyBwgIEAAYgAQyBwgJEAAYgATSAQkxNjI4NmowajSoAgewAgE&sourceid=chrome&ie=UTF-8"


func main(){
	fmt.Print("get request")
	// here variable response it will hold  GET http response from that  which it was reuslt of get request from that cite of page  
	response,err := http.Get(url)
	if err != nil{
		panic(err)

	}
	// checking for data type that return as  response 
	fmt.Printf("request is %T:", response)
	// we must close the request to avoid issue of data leak and also we close to have good memmory management
	defer response.Body.Close()
	// response.bodyb contain the body data incluade the heading , title and another  element from the response 
	response2,err := io.ReadAll(response.Body)

	if err != nil{
		panic(err)
	}
	fmt.Print("content is:",string(response2))

    ex := ExpoPushMessage{
		Body: "hello world",
		To: "issa mwanga",
		Title: "greeting",
		Sound: "default",
	}
	printJson(ex)
    

}
func printJson(ex ExpoPushMessage){
	payload , _ := json.Marshal([]ExpoPushMessage{ex})
	fmt.Println(payload)
}

