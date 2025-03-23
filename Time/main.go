package main

import (
	"fmt"
	"time"
)

func main() {
	// create current time or   present time we use Now to method 
	presenttime := time.Now()
	fmt.Println(presenttime);
	// to make it in good format we use Formart method  and inorder to find that it must be written in this format  as golang suggest   01-02-2006 Monday 15:04:05
	fmt.Println(presenttime.Format("01-02-2006 Monday"))

	// to create date   
	createdate := time.Date(2002,time.October,07,23,02,23,0,time.UTC)
	fmt.Println(createdate.Format("01-02-2006 15:04:05 Monday"))
    saido :=time.Date(2024,time.February,05,02,34,34,12,time.UTC)
	fmt.Println(saido.Format("01-02-2006 15:04:05 Monday"))

}