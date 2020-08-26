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

type shape interface {
	area() float64
	perimeter() float64
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

func print(s shape) {
	fmt.Printf("Shape: %#v \n", s)
	fmt.Printf("Area:  %#v \n", s.area())
	fmt.Printf("Perimeter:  %#v \n", s.perimeter())
}
func main() {
	var s shape
	fmt.Printf("%T\n", s) //<nil>

	ball := circle{radius: 2.5}
	s = ball

	print(s)
	fmt.Printf("Type of s: %T\n", s)
	// 	Shape: main.circle{radius:2.5}
	// Area:  19.634954084936208
	// Perimeter:  15.707963267948966
	// Type of s: main.circle

	room := rectangle{width: 2, height: 3}
	s = room
	fmt.Printf("Type of s: %T\n", s)//Type of s: main.rectangle
}
