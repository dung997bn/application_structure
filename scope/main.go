package main

import (
	"fmt"
	f "fmt"
)

var u = 18

const done = false //package scope
func main() {
	//f.Println("ok \n")
	f1()
	var task = "Runing" // local(block) scope
	f.Println(task, done)

	done := true
	fmt.Printf("done in main: %v \n", done)
	fmt.Println("Bye Bye")
}

func f1() {
	fmt.Printf("done in f1(): %v \n", done)
}
