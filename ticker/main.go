package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			PrintTime()
		}
	}

}

func PrintTime() {
	fmt.Println("Running server at ", time.Now().Format("15:04:05"))
}
