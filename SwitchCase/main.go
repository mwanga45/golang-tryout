package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())  // Seed the random generator
    diceNumber := rand.Intn(6) + 1    // Generate a random number between 1 and 6
    fmt.Println("Dice roll:", diceNumber)
}
