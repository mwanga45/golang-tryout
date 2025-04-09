package main

import (
	"fmt"
	"time"
)

func main() {
	time.AfterFunc( 1 *time.Minute, func() {
		result := Printlns(5,4)
		fmt.Println(result)
	})
	time.Sleep(1500 *time.Second)
}

func Printlns(x int, y int) int {
	Result := x + y
	return Result
}