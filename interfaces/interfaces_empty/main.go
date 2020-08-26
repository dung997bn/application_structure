package main

import "fmt"

type emptyInterface interface {
}

type person struct {
	info interface{}
}

func main() {
	var empty interface{}
	empty = 5
	fmt.Println(empty) //5

	empty = "GO"
	fmt.Println(empty) //GO

	empty = []int{5, 6, 9}
	fmt.Println(empty) //[5 6 9]
	//fmt.Println(len(empty))//invalid argument empty (type interface {}) for len
	fmt.Println(len(empty.([]int))) //3

	you := person{}
	you.info = "Your name"
	fmt.Println(you.info) //Your name

	//empty interface can store any type

	you.info = 3
	fmt.Println(you.info) //3

	you.info = []float64{4., 6., 8.9}
	fmt.Println(you.info)//[4 6 8.9]

}
