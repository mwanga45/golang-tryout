package main

import (
	"fmt"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request){
	ctx := r.Context()
	fmt.Println("Hello wrold")

	select{
	case <- time.After(10 * time.Second):
		fmt.Fprintf(w,"Successfuly return")
	case <- ctx.Done():
		err :=  ctx.Err()
		fmt.Println("Something wrong", err)
		internalError := http.StatusInternalServerError
		http.Error(w, err.Error(), internalError)
		
		
	}
}
func main() {
	http.HandleFunc("/hello", Hello)
	http.ListenAndServe(":2000", nil)

}