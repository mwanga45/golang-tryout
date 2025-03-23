package main

import "fmt"

func main() {
	// struct is  composite data  that group  together field under a single name  that may allowed you to create more complex data
	//the name of struct must h start with capital letter
	info1 := Student{"issa", "Mwanga", true, 16}
	fmt.Printf("what is your Firstname and lastname: \n firstname %v  and lastname %v",info1.Firstname, info1.Lastname)

	// also u can print inter struct data 
     fmt.Println(info1)

}

type Student struct {
	Firstname string
	Lastname  string
	Register  bool
	Age       int
}