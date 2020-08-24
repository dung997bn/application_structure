package main

import "fmt"

func f1() {
	fmt.Println("This is funtion f1 in go")
}

func f2(a int, b string) {
	fmt.Println(a, b)
}

func f3(a, b, c int, d float64) {
	fmt.Println(a, b, c)
}

func f4(a float64, b float64) float64 {
	return a + b
}

func f5(a, b int, c string) (int, string) {
	return a - b, "Ok " + c
}

func f6(a, b int) (sum int) {
	fmt.Println(sum) //default 0
	sum = a + b
	fmt.Println(sum)
	return sum
}
func main() {
	f1()            //This is funtion f1 in go
	f2(2, "Ok")     //2 Ok
	f3(1, 3, 5, 56) //1 3 5
	s := f4(4, 5)
	fmt.Println(s) //9

	d, p := f5(10, 3, "halo")
	fmt.Println(d, p) //7 Ok halo

	f6(10, 29)
	// 	0
	// 39
}
