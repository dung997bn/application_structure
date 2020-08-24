package main

import "fmt"

func main() {
	name := "Adrei"
	fmt.Println(&name) //0xc00003e1f0

	var x int = 2
	ptr := &x

	fmt.Printf("ptr is of type %T with a value of %v\n", ptr, ptr) //ptr is of type *int with a value of 0xc0000a2058
	fmt.Printf("address of x is %p \n", &x)                        //address of x is 0xc0000120b0

	var ptr1 *float64 //zero value is nil
	_ = ptr1
	p := new(int)
	x = 100
	p = &x
	fmt.Printf("p is of type %T with a value of %v \n", p, p) //p is of type *int with a value of 0xc0000120b0
	fmt.Printf("address of x is %p \n", &x)                   //address of x is 0xc0000120b0

	*p = 90                        // equal to x=90
	fmt.Println(x, *p)             //90 90
	fmt.Println("*p==x:", *p == x) //*p==x: true

	*p = 10        //x=10
	*p = *p / 2    //mean x/2
	fmt.Println(x) //5

	//&value =>pointer
	// *pointer =>value

}
