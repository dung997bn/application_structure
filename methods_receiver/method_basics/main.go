package main

import (
	"fmt"
	"time"
)

type names []string

//receiver method
func (n names) print() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day) //time.Duration

	seconds := day.Seconds()
	fmt.Printf("%T\n", seconds)                   //float64
	fmt.Printf("Seconds in a day: %v\n", seconds) //Seconds in a day: 86400

	friends := names{"Dan", "Mary"}
	friends.print()
	// 	0 Dan
	// 1 Mary

	var n int64 = 943535
	fmt.Println(n)                //943535
	fmt.Println(time.Duration(n)) //943.535µs
}
