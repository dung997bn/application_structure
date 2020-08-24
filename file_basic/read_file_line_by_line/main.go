package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("my_file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	// 	First line found:  ok
	// ok
	// halao
	// vl

	scanner.Split(bufio.ScanRunes)

	// 	First line found:  o
	// k

	// o
	// k

	// h
	// a
	// l
	// a
	// o

	// v
	// l
	success := scanner.Scan()
	if success == false {
		err = scanner.Err()
		if err == nil {
			log.Println("Scan was completed and it reached EOF")
		} else {
			log.Fatal(err)
		}
	}

	fmt.Println("First line found: ", scanner.Text()) //First line found:  ok ok

	for scanner.Scan() {
		fmt.Println(scanner.Text())
		//halao

		//vl
	}

	if err := scanner.Err(); err != nil {
		log.Fatal()
	}
}
