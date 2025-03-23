package main

import (
	"fmt"
	"unicode/utf8"
	"github.com/boltdb/bolt"
)

func main() {
	// arrey initilization
	// note that during initilization of an array we should we put number of items
	// that goes inside it eg [6] means there is six items will go inside the arreys for this format bellow only
	var balls [6]string

	balls[0] = "Basketball"
	balls[1] = "vollyball"
	balls[2] = "football"
	balls[3] = "America-football"
	// or you can create [] like
	sports := []string{"football", "baskateball", "tennes", "Rugby"}

	fmt.Println("The values of An Arreys is :", balls)
	fmt.Println("The values of An Arreys is :", sports)
	fmt.Println("The values of An Arreys in second index is :", balls[2])
	fmt.Println("The values of An Arreys is :", len(balls[2]))
	fmt.Println("The values of An Arreys is :", len(balls))
	fmt.Println("The values of An Arreys is :", utf8.RuneCountInString(balls[3]))
	// the space lived in the  up  first output indicate that there is 2 empty index value that have null value (there value are empty string)

	var test = []string{"database"}
	fmt.Print(test,"this is test")
}
