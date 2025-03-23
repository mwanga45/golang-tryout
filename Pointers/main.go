package main

import "fmt"

func main() {
	// var mypointer *int  
	// fmt.Println("value of Pointer is " , mypointer)

	// number := 24;
	// //this will line it reference the  memory location 
	// mpointer := &number
    // // this will show the memory location where you can get actual value
	// fmt.Println("value of memory location is:", mpointer)
	// //this  will be able to find the actual value  which pointer contain
	// fmt.Println("value of memory location is:", *mpointer)
	// // to make sure this pointer is been able to change the actual value in memory location here is example of it 
	// *mpointer = *mpointer / 2

	// fmt.Println("the new value of number is :", number)
	// // throught pointer  u will beable to update or change the actual value and through this tricks you will avoid creating of copy eg
	
	// number2 := number 
	//  fmt.Println("value of number2  is :", number2)
	//  number2 = 2
	//  fmt.Printf("see the value of number doesnt change even if i was already say number2 = number \n and then change the value of number2 and say equal to 12  it will remain to (%v) because there iwas just make copy not change the actual value of number  : ",number)
	prt := 23

	mypointer := &prt
	fmt.Println("value of the memory reference is ", mypointer)
	fmt.Println("The actual value inside that memory reference is ", *mypointer)
    *mypointer = *mypointer + 2
	fmt.Println("Update value is ", prt)
}