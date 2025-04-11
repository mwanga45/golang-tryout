package main

import (
	"fmt"
	"time"
)

func main() {
	time.AfterFunc(30 *time.Second, func() {
		result := Additional(5, 4)
		fmt.Println("The answer is:",result)

	})
	time.Sleep(45 *time.Second)

}

func Additional(x int, y int) int {
	Result := x + y
	return Result
}
 