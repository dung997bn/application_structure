package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byetSlice := make([]byte, 2)

	numberBytesRead, err := io.ReadFull(file, byetSlice)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("number of bytes read: %d \n", numberBytesRead)
	log.Printf("Data read: %s \n", byetSlice)

	// file, err = os.Open("main.go")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// data, err := ioutil.ReadAll(file)

	// fmt.Printf("Data as string: %s\n", data)
	// fmt.Println("number of bytes read:", len(data))

	data, err := ioutil.ReadFile("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Data read: %s\n", data) //Data read: Hello Kirin
}
