package main

import "fmt"

func main() {
	// 	s1 := "Hello Kirin"
	// 	fmt.Println(s1)

	// 	fmt.Println("He says \"Ok\"") //He says "Ok"
	// 	fmt.Println(`He says "Ok"`)   //He says "Ok" this call raw string

	// 	s2 := "Price: 1000 \nBrand: Nike"
	// 	fmt.Println(s2)
	// 	fmt.Println(`
	// Price: 1000
	// Brand: Nike`)

	// 	fmt.Println(`C:\User\Document`)   //C:\User\Document
	// 	fmt.Println("C:\\User\\Document") //C:\User\Document

	var s3 = "I love " + "Go " + "Programming"
	fmt.Println(s3 + "!")   //I love Go Programming!
	fmt.Printf("%s \n", s3) //I love Go Programming
	fmt.Printf("%q \n", s3) //"I love Go Programming"
}
