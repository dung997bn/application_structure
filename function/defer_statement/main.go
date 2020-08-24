package main

import (
	"fmt"
	"log"
	"os"
)

func foo() {
	fmt.Println("This is foo")
}
func bar() {
	fmt.Println("This is bar")
}

func foobar() {
	fmt.Println("This is foobar")
}

func main() {
	// foo() //This is foo

	// bar() //This is bar

	defer foo() // defer will excute the function at the end of progams: First in last out
	bar()

	fmt.Println("Just a string after defering foo and calling bar")
	// 	This is bar
	// Just a string after defering foo and calling bar
	// This is foo

	defer foobar()
	// 	This is bar
	// Just a string after defering foo and calling bar
	// This is foobar
	// This is foo

	file, err := os.Open("main.go")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}
