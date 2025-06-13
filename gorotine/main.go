package main

import (
	"fmt"
	"sync"
	// "time"
)

func printnames(name string, count int, wg *sync.WaitGroup) {
     defer  wg.Done()
	 for i := 0 ; i < count; i++{
		fmt.Printf("%s %d\n", name,i)
		// time.Sleep(100 * time.Millisecond)
		fmt.Println()
	 }
}

func main() {
     var wg  sync.WaitGroup
	 wg.Add(2)
	 go printnames("Issa", 5, &wg)
	 go printnames("Tach", 10, &wg)

	 fmt.Println("Begin rotine")

	 wg.Wait()
	 fmt.Println("Successfuly complite")
}