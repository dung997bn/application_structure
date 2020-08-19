package main

import "fmt"

func main() {
	// var numbers [4]int

	// fmt.Printf("%v \n", numbers)

	// fmt.Printf("%#v \n", numbers)

	// var a1 = [4]float64{}
	// fmt.Printf("%#v \n", a1)

	// var a2 = [3]float64{-10, 2, 90}
	// fmt.Printf("%#v \n", a2)

	// a3 := [4]string{"Dung", "Dao", "Ha", "Quynh"}
	// fmt.Printf("%#v \n", a3)

	// a4 := [4]string{"Ha"}

	// fmt.Printf("%#v \n", a4)

	// //ellipsis operator
	// a5 := [...]int{-12, -3, -5, 5, 6, 5, 7, 78}
	// fmt.Printf("%#v \n", a5)
	// fmt.Printf("len of a5: %d \n", len(a5))

	// numbers := [3]int{1, 2, 4}
	// fmt.Printf("%#v \n", numbers)

	// numbers[1] = 3
	// fmt.Printf("%#v \n", numbers)

	//two way
	// for i, v := range numbers {
	// 	fmt.Println("index: ", i, " value: %d", v)
	// }

	//fmt.Println(strings.Repeat("#", 20)) //####################
	// for i := 0; i < len(numbers); i++ {
	// 	fmt.Println("index: ", i, " value: %d", numbers[i])
	// }

	// balances := [2][3]int{
	// 	{5, 6, 7},
	// 	[3]int{8, 9, 10},
	// }
	// fmt.Println(balances)

	// m := [3]int{4, 7, 9}
	// n := m
	// fmt.Println("n is equal to m: ", n == m)//true

	// grades := [3]int{
	// 	1: 10,
	// 	0: 5,
	// 	2: 7,
	// }

	// fmt.Println(grades)//[5 10 7]

	// accounts := [3]int{2: 50}

	// fmt.Println(accounts)//[0 0 50]

	// names := [...]string{
	// 	5: "Dung",
	// }
	// fmt.Println(names)//[     Dung] len(6)

	// cities := [...]string{
	// 	5:        "paris",
	// 	"london", //no key: last element
	// 	1:        "Hanoi",
	// }

	// fmt.Println(cities)
	// fmt.Printf("%#v \n", cities)//[7]string{"", "Hanoi", "", "", "", "paris", "london"}

	weekend := [7]bool{5: true, true}
	fmt.Println(weekend)

}
