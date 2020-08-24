package main

import "fmt"

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}

func main() {
	func(msg string) {
		fmt.Println(msg)
	}("I am anonymous function")

	//I am anonymous function

	a := increment(9)
	fmt.Printf("%T \n", a) //func() int

	fmt.Println(a()) //10

	fmt.Println(a()) //11
	a()
	fmt.Println(a()) //13
}
