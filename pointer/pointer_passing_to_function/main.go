package main

import "fmt"

func change(a *int) *float64 {
	*a = 100
	b := 5.5
	return &b
}

func changeVar(a int) {
	a = 66
}
func main() {
	x := 8
	p := &x
	fmt.Println("Value of x before calling change(): ", x) //8
	change(p)
	fmt.Println("Value of x after calling change(): ", x) //100

	fmt.Println("Value of x before calling changeVar(): ", x) //100
	changeVar(x)
	fmt.Println("Value of x after calling changeVar(): ", x) //100
}
