
package main

import (
	"fmt"
	"math"
)

// Define an interface type that will hold method signatures.
type Shape interface {
	Area() float64
	Para() float64
}

// Define the `circle` struct type for circle-related calculations.
type circle struct {
	radius float64
}

// Implement the `Area` method for `circle`.
func (c circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

// Implement the `Para` method for `circle` (calculates circumference).
func (c circle) Para() float64 {
	return 2 * math.Pi * c.radius *c.radius
}

// Define the `Rectangle` struct type for rectangle-related calculations.
type Rectangle struct {
	width, height float64
}

// Implement the `Area` method for `Rectangle`.
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// Implement the `Para` method for `Rectangle` (calculates perimeter).
func (r Rectangle) Para() float64 {
	return 2*r.width + 2*r.height
}

// Function to measure and display the properties of a `Shape`.
func measure(g Shape) {
	fmt.Println("Shape:", g)
	fmt.Println("Area:", g.Area())
	fmt.Println("Perimeter:", g.Para())
}

func main() {
	r := Rectangle{height: 3, width: 4} // Create a Rectangle instance.
	c := circle{radius: 2}              // Create a circle instance.

	measure(r) // Measure and display properties of Rectangle.
	measure(c) // Measure and display properties of circle.
}

