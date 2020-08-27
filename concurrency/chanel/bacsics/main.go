package main

import "fmt"

func main() {
	var ch chan int

	fmt.Println(ch) //<nil>

	ch = make(chan int)
	fmt.Println(ch)

	c := make(chan int)
	//<-chanel operator

	//Send
	c <- 10

	//Receive
	num := <-c
	fmt.Println(<-c)
	_ = num
	close(c)
}
