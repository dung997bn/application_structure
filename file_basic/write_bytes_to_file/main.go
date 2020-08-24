package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"b.txt",
		os.O_WRONLY|os.O_TRUNC|os.O_CREATE,
		0644,
	)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	byteSlices := []byte("I learn golang")
	byteWritten, err := file.Write(byteSlices)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes Written: %d\n", byteWritten)

	bs := []byte("Go programming is cool!")
	err = ioutil.WriteFile("c.txt", bs, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
