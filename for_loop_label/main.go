package main

import "fmt"

func main() {
	//label
	// 	people := [5]string{"Dung", "Quynh", "Dao", "Bich", "Ha"}
	// 	friends := [2]string{"Trung", "Quynh"}

	// outer:
	// 	for index, name := range people {
	// 		for _, friend := range friends {
	// 			if name == friend {
	// 				fmt.Printf("Found a friend %q at index %d \n", friend, index)
	// 				break outer
	// 			}
	// 		}
	// 	}

	// 	fmt.Println("Next instruction after the break")

	//go to
	i := 0
loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop
	}

	//error
	// 	goto todo
	// 	x := 5
	// todo:
	// 	fmt.Println("somthing here")
}
