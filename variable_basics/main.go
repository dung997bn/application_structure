package main

import "fmt"

func main() {
	car, cost := "Audi", 5000
	fmt.Println(car, cost)
	car, year := "BMW", 2019
	_ = year

	var open = false
	open, file := true, "a.txt"
	_, _ = open, file

	var (
		salary    float64
		firstName string
		gender    bool
	)

	fmt.Println(salary, firstName, gender)

	// var a, b, c int
	// fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8
	_, _ = i, j

	j, i = i, j //swapping variables
	fmt.Println(i, j)

	sum := i + j
	fmt.Println(sum)

	var a = 5
	var b = 7.8

	a = int(b)
	fmt.Println(a, b)

	// var x int
	// x = int("5")
	// var value int
	// var price float64
	// var name string
	// var done bool

	// fmt.Println(value, price, name, done)

	// var Hoooooo_Kiiiiiiiii int = 10
	// fmt.Println(Hoooooo_Kiiiiiiiii)

	fmt.Printf("Your age is %f", 20.2)
	fmt.Println("He saids \"HELLO\"")

	var pi = 3.14
	var shape = "Circle"
	var month = 12

	//%v---> replace any type
	fmt.Printf("Pi is %v \t Share is %v \t Month is %v \n", shape, pi, month)

	//%T---> type
	fmt.Printf("Pi is %T \t Share is %v \t Month is %v \n", shape, pi, month)

	//%t ->bool
	closed := true
	fmt.Printf("File closed: %t \n", closed)

	//%b ->base 2
	fmt.Printf("%b \n", 55)

	x := 3.4
	y := 6.9
	fmt.Printf("x*y=%.2f \n", x*y)

}
