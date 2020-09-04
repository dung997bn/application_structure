package main

import (
	"fmt"

	calcv1 "github.com/dung997bn/go_math/calc"
	"github.com/dung997bn/go_math/geometry"
	calcv2 "github.com/dung997bn/go_math/v2/calc"
)

func main() {
	fmt.Println(geometry.CubeVolumne(0))

	z := calcv1.Add(1, 2)
	fmt.Println(z)
	s := calcv2.Add(5, 6, 7, 8, 9)
	fmt.Println(s)
}
