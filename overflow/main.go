package main

import (
	"fmt"
	"math"
)

func main() {
	// var x uint8 = 255 //x=0 because uint8
	// x++               //overflow
	// fmt.Println(x)    //x=0

	// a := int8(255 + 1//error)

	var b int8 = 127   // int8 : -128-->127
	fmt.Println(b + 1) // -128

	var c int8 = -128
	c -= 6
	fmt.Println(c) // -122

	f := float32(math.MaxFloat32)
	fmt.Println(f)
}
