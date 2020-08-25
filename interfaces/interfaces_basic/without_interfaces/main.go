package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	width, height float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2) //pi *r binh
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) area() float64 {
	return (r.width + r.height) / 2
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

func printCircle(c circle) {
	fmt.Println("Shape", c)
	fmt.Println("Area:", c.area())
	fmt.Println("Perimeter: ", c.perimeter())
}

func printRectangle(r rectangle) {
	fmt.Println("Shape", r)
	fmt.Println("Area:", r.area())
	fmt.Println("Perimeter: ", r.perimeter())
}

func main() {
	c1 := circle{radius: 5}
	r1 := rectangle{width: 50.4, height: 23.1}

	printCircle(c1)
	// 	Shape {5}
	// Area: 78.53981633974483
	// Perimeter:  31.41592653589793
	printRectangle(r1)
	// 	Shape {50.4 23.1}
	// Area: 36.75
	// Perimeter:  147
}
