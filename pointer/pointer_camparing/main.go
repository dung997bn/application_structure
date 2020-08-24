package main

import "fmt"

func main() {
	a := 5.5
	p1 := &a
	pp1 := &p1
	fmt.Printf("Value of p1: %v, address of p1: %v \n", p1, &p1)     //Value of p1: 0xc0000a2058, address of p1: 0xc0000cc018
	fmt.Printf("Value of pp1: %v, address of pp1: %v \n", pp1, &pp1) //Value of pp1: 0xc0000cc018, address of pp1: 0xc0000cc020

	fmt.Printf("*p1 is %v\n", *p1)   //*p1 is 5.5
	fmt.Printf("*pp1 is %v\n", *pp1) //*pp1 is 0xc000012090

	fmt.Printf("**pp1 is %v\n", **pp1) //**pp1 is 5.5

	**pp1++
	fmt.Printf("a is %v\n", a) //a is 6.5

	//**COMPARING POINTERS**//

	var p2 *int
	fmt.Printf("%#v \n", p2) //(*int)(nil)

	y := 5
	p2 = &y
	z := 5
	p3 := &z
	fmt.Println(p2 == p3) //false

	p4 := &y
	fmt.Println(p2 == p4) //true
}
