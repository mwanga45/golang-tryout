package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Please rating me my speed of coding up to 5: ")

	//comma ok || error comma
	input1, _ := reader.ReadString('\n')
	// input1 = strings.TrimSpace(input1)
	input,errs := strconv.ParseFloat(input1,32)
	if errs !=nil{
		fmt.Println("failed to convert the number", errs)
	}
	if input == 2{
		fmt.Println("ooH that is moderate",input1)
	}else if input ==4||input == 3{
		fmt.Println("Oooh that is so great soon i will reach 5 ", input)
	}else if input == 0||input==1{
		fmt.Println("ooh that is so bad in neeed to Improve",input)
	}else if input > 5{
		fmt.Println("Please read the message provided carefuly",input)
	}else{
		fmt.Println("Yeah Icant beleive you rating me that match wow", input)
	}
   reader2 := bufio.NewReader(os.Stdin)
   fmt.Println("Please rate me:")
   meg,_ := reader2.ReadString('\n')
   input2,err := strconv.ParseFloat(strings.TrimSpace(meg),32)

   if err !=nil{
	fmt.Println("something goes wrong", err)
   }else{
	fmt.Println("Okey thanks for rating:",input)
   }
   numbrating := input2 + 1
   fmt.Println("rating after add one per Each rate:", numbrating)

	
}