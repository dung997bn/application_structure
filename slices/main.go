package main

func main() {
	// var cities []string
	// fmt.Println("cities is equal to nil: ", cities == nil) //true
	// fmt.Printf("cities %#v \n", cities)
	// fmt.Println(len(cities)) //0

	// cities[0] = "London" //runtime error: index out of range [0] with length 0
	// fmt.Println(len(cities))
	//numbers := []int{3, 434, 35, 54, 7}
	// fmt.Println(numbers) //[3 434 35 54 7]

	// nums := make([]int, 2)
	// fmt.Printf("%#v \n", nums) //[]int{0, 0}

	// type names []string
	// friends := names{"Dung", "Nga"}
	// fmt.Println(friends)

	// myFriend := friends[0]
	// fmt.Println("My best friend is ", myFriend) //My best friend is  Dung

	// friends[0] = "Kirin"
	// fmt.Println("My best friend is ", myFriend) //My best friend is  Dung

	// for index, value := range numbers {
	// 	fmt.Printf("index : %v, value: %v \n", index, value)
	// }

	// var n []int
	// n = numbers
	// fmt.Println(n)

	// var n []int
	// fmt.Println(n == nil) //true

	// m := []int{}
	// fmt.Println(m == nil) //false

	//compare 2 slices
	// a, b := []int{1, 2, 3}, []int{1, 2, 5, 3}
	// //fmt.Println(a == b) //error invalid operation: a == b (slice can only be compared to nil)

	// var eq bool = true
	// for i, valueA := range a {
	// 	if valueA != b[i] {
	// 		eq = false
	// 		break
	// 	}
	// }
	// if len(a) != len(b) {
	// 	eq = false
	// }
	// if eq {
	// 	fmt.Println("Equal")
	// } else {
	// 	fmt.Println("Not Equal")
	// }

	// numbers := []int{1, 2, 4}
	// numbers = append(numbers, 10)
	// fmt.Println(numbers) //[1 2 4 10]

	// numbers = append(numbers, 20, 30, 1444)
	// fmt.Println(numbers) //[1 2 4 10 20 30 1444]

	// n := []int{100, 200}
	// numbers = append(numbers, n...)
	// fmt.Println(numbers) //[1 2 4 10 20 30 1444 100 200]]

	// src := []int{10, 20, 30, 35} //[10 20 30 35]
	// dst := make([]int, 2)        //[10 20]
	// nn := copy(dst, src)         //2
	// fmt.Println(src, dst, nn)

	//Slices Expression
	// a := [5]int{1, 2, 4, 5, 6}

	// //a[start:stop]
	// b := a[1:3]
	// fmt.Printf("%v , %T", b, b) //[2 4] , []int

	// s1 := []int{1, 2, 3, 4, 5, 6, 7}
	// s2 := s1[0:6]
	// fmt.Println(s2)
	// s3 := s1[2:]    //same as s1[2:len(s1)]
	// fmt.Println(s3) //[3 4 5 6 7]
	// s4 := s1[:2]    //same as s1[0:2]
	// s5 := s1[:]     //same as s1[2:len(s1)]
	// fmt.Println(s4) //[1 2]
	// fmt.Println(s5) //[1 2 3 4 5 6 7]
	// //s4 := s1[:100]//error
	// s7 := append(s1[:4], 177) //[1 2 3 4 177]
	// fmt.Println(s7)
}
