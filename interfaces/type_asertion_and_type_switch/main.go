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

func (c circle) volumn() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func main() {
	var s shape = circle{radius: 2.5}
	fmt.Printf("%T\n", s)

	//s.volumn()//error :s.volumn undefined (type shape has no field or method volumn)
	//s.(circle).volumn()//no error

	//type asertion
	ball, ok := s.(circle)
	if ok == true {
		fmt.Printf("Ball volume : %v\n", ball.volumn()) //Ball volume : 49.087385212340514
	}

	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value) //main.circle{radius:2.5} has circle type
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)
	}

	s = rectangle{height: 5, width: 6}
	switch value := s.(type) {
	case circle:
		fmt.Printf("%#v has circle type\n", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type\n", value)//main.rectangle{width:6, height:5} has rectangle type
	}
}
