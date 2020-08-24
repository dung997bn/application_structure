package main

import (
	"fmt"
	"strings"
)

func f1(a ...int) {
	fmt.Printf("%T \n", a)  //[]int
	fmt.Printf("%#v \n", a) //[]int{1, 232, 3, 435, 67, 8}
}

func f2(a ...int) {
	a[0] = 90
}

//sum
func SumAndProduct(a ...float64) (float64, float64) {
	sum := 0.
	product := 1.

	for _, v := range a {
		sum += v
		product *= v
	}

	return sum, product
}

func personInfomation(age int, name ...string) string {
	fullName := strings.Join(name, "-")
	returnString := fmt.Sprintf("Age: %d, FullName: %s", age, fullName)
	return returnString
}
func main() {
	// f1(1, 232, 3, 435, 67, 8)
	// f1()
	// // []int
	// // []int(nil)

	nums := []int{1, 2}
	nums = append(nums, 4, 6)
	// f1(nums...)
	// []int
	// []int{1, 2, 4, 6}

	// f2(nums...)
	// fmt.Println("Nums: ", nums)//Nums:  [90 2 4 6]

	s, p := SumAndProduct(2.0, 4.5, 9.5)
	fmt.Println(s, p) //16 85.5

	info := personInfomation(50, "halo", "Kirrin", "DUng")

	fmt.Println(info) //Age: 50, FullName: halo-Kirrin-DUng

	var name = []string{"Quara", "Penta"}
	info1 := personInfomation(50, name...)

	fmt.Println(info1) //Age: 50, FullName: Quara-Penta

}
