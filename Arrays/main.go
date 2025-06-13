package main

import (
	"fmt"
	"time"
	// "unicode/utf8"

)

func main() {
	// arrey initilization
	// note that during initilization of an array we should we put number of items
	// that goes inside it eg [6] means there is six items will go inside the arreys for this format bellow only
	// var balls [6]string

	// balls[0] = "Basketball"
	// balls[1] = "vollyball"
	// balls[2] = "football"
	// balls[3] = "America-football"
	// // or you can create [] like
	// sports := []string{"football", "baskateball", "tennes", "Rugby"}

	// fmt.Println("The values of An Arreys is :", balls)
	// fmt.Println("The values of An Arreys is :", sports)
	// fmt.Println("The values of An Arreys in second index is :", balls[2])
	// fmt.Println("The values of An Arreys is :", len(balls[2]))
	// fmt.Println("The values of An Arreys is :", len(balls))
	// fmt.Println("The values of An Arreys is :", utf8.RuneCountInString(balls[3]))
	// // the space lived in the  up  first output indicate that there is 2 empty index value that have null value (there value are empty string)

	// var test = []string{"database"}
	// fmt.Print(test,"this is test")

  allocatedtae := make(map[int]string, 4)
  now := time.Now()
  for i:= 0;i <4 ; i++{
     d :=  now.Add(time.Duration(i) * 24 * time.Hour)
	 allocatedtae[int(d.Weekday())] =  d.Format("2006-1-02")
	 

  }
  fmt.Println(allocatedtae)
}
