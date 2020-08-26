package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}
type object interface {
	volumn() float64
}

//geometry is embedding shape and object interfaces
type geometry interface {
	shape
	object
	getColor() string
}

type cube struct {
	edge  float64
	color string
}

func (c cube) area() float64 {
	return 6 * (c.edge * c.edge)
}

func (c cube) volumn() float64 {
	return math.Pow(c.edge, 3)
}

func (c cube) getColor() string {
	return c.color
}

func measure(g geometry) (float64, float64) {
	a := g.area()
	v := g.volumn()
	return a, v
}

func main() {
	c := cube{edge: 2}
	a, v := measure(c)
	fmt.Printf("Area: %v, Volumn: %v \n", a, v) //Area: 24, Volumn: 8
}
