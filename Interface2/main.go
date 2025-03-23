package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Parameter() float64
	Area() float64
}
type Reactangle struct {
	width  float64
	height float64
}


func (r Reactangle) Parameter() float64 {
	return r.width + r.height
}
func (r Reactangle) Area() float64{
	return r.width * r.height
	
}
type Circle struct {
	radius float64
}


func (c Circle) Area() float64 {
	return math.Pi  * c.radius * c.radius
 }
 func (c Circle)Parameter() float64 {
	return  2* c.radius * c.radius
	
 }

func message(msg Shape) {
	fmt.Println("The interface is", msg)
	fmt.Println("The Parameter of the reactangle is: ", msg.Parameter())
	fmt.Println("Thee Area is of rectangle is :", msg.Area())

}
func main() {
	r := Reactangle{height: 3, width: 2}
	c := Circle{radius: 4}

	message(r)
	message(c)

}
