package main

import "fmt"

type km float64
type mile float64

func main() {
	// type speed uint
	// var s1 speed = 10
	// var s2 speed = 20
	// fmt.Println(s2 - s1) //10

	// var x uint
	// //x = s1 //error because diffirent type
	// x = uint(s1)

	// fmt.Println(x)

	// var s3 speed = speed(x) //convert

	// fmt.Println(s3)

	// var parisToLondon km = 2321312
	// var distanceInMiles mile

	// //distanceInMiles = parisToLondon / 0.621//error

	// distanceInMiles = mile(parisToLondon / km(0.621))//true
	// fmt.Println(distanceInMiles)

	var a uint8 = 10
	var b byte

	b = a
	_ = b

	type second uint
	var hour second = 3600
	fmt.Printf("Minutes in an hour : %d \n", hour/second(a))
}
