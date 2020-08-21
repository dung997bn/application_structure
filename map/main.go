package main

import "fmt"

func main() {
	// var employees map[string]string
	// fmt.Printf("%#v\n", employees) //map[string]string(nil)

	// fmt.Printf("No of pairs:%d\n", len(employees)) //No of pairs:0

	// fmt.Printf("The value for key %q is %q \n", "Dan", employees["Dan"]) //The value for key "Dan" is ""

	// var accounts map[string]float64
	// fmt.Printf("%#v\n", accounts["0x323"]) //0

	// var m1 map[[5]int]string
	// _ = m1
	// employees["Dan"] = "Programmer" //error: assignment to entry in nil map

	// people := map[string]float64{}
	// people["John"] = 24
	// people["Marry"] = 67.3
	// fmt.Println(people) //map[John:24 Marry:67.3]

	// map1 := make(map[int]int)
	// map1[4] = 8
	// fmt.Println(map1) //map[4:8]

	// balances := map[string]float64{
	// 	"USD": 323.3,
	// 	"VND": 2333.5,
	// 	//50: 44304.6 error , all keys the same type
	// 	"CFH": 3545,
	// }
	// fmt.Println(balances) //map[CFH:3545 USD:323.3 VND:2333.5]

	// m := map[int]int{1: 10, 2: 20, 3: 30}
	// fmt.Println(m) //map[1:10 2:20 3:30]

	// balances["USD"] = 112.3
	// fmt.Println(balances) //map[CFH:3545 USD:112.3 VND:2333.5]

	// balances["GPD"] = 2232
	// fmt.Println(balances) //map[CFH:3545 GPD:2232 USD:112.3 VND:2333.5] not update if no key

	// fmt.Println(balances["HYI"]) //0 no key

	// // v, ok := balances["HYI"]
	// // if ok {
	// // 	fmt.Println("The HYI balance is :", v)
	// // } else {
	// // 	fmt.Println("The HYI doesn't exist")
	// // } //The HYI doesn't exist

	// balances["HYI"] = 123.5
	// v, ok := balances["HYI"]
	// if ok {
	// 	fmt.Println("The HYI balance is :", v)
	// } else {
	// 	fmt.Println("The HYI doesn't exist")
	// } //The HYI balance is : 123.5

	// // for k, v := range balances {
	// // 	fmt.Printf("Key: %#v, Value: %#v \n", k, v)
	// // }
	// // Key: "USD", Value: 112.3
	// // Key: "VND", Value: 2333.5
	// // Key: "CFH", Value: 3545
	// // Key: "GPD", Value: 2232
	// // Key: "HYI", Value: 123.5

	// //delete
	// delete(balances, "CFH")
	// for k, v := range balances {
	// 	fmt.Printf("Key: %#v, Value: %#v \n", k, v)
	// }
	// 	Key: "GPD", Value: 2232
	// Key: "HYI", Value: 123.5
	// Key: "USD", Value: 112.3
	// Key: "VND", Value: 2333.5

	//Compare
	// a := map[string]string{"A": "X"}
	// b := map[string]string{"B": "Y"}
	// //fmt.Println(a == b)// "invalid operation: a == b (map can only be compared to nil)",

	// s1 := fmt.Sprintf("%s", a)
	// s2 := fmt.Sprintf("%s", b)

	// fmt.Println(s1) //map[A:X]
	// fmt.Println(s2) //map[B:Y]
	// if s1 == s2 {
	// 	fmt.Println("Maps are equal")
	// } else {
	// 	fmt.Println("Maps are not equal")
	// }//Maps are not equal

	// a := map[string]string{"A": "U"}
	// b := map[string]string{"A": "U"}
	// //fmt.Println(a == b)// "invalid operation: a == b (map can only be compared to nil)",

	// s1 := fmt.Sprintf("%s", a)
	// s2 := fmt.Sprintf("%s", b)

	// fmt.Println(s1) //map[A:U]
	// fmt.Println(s2) //map[A:U]
	// if s1 == s2 {
	// 	fmt.Println("Maps are equal")
	// } else {
	// 	fmt.Println("Maps are not equal")
	// } //Maps are equal

	friends := map[string]int{"Dung": 24, "Nga": 28}
	neighbors := friends
	friends["Dung"] = 25

	fmt.Println(friends)   //map[Dung:25 Nga:28]
	fmt.Println(neighbors) //map[Dung:25 Nga:28]

	people := make(map[string]int)

	for k, v := range friends {
		people[k] = v
	}
	people["Dung"] = 27
	people["An"] = 30
	fmt.Println(friends)
	fmt.Println(people) //map[An:30 Dung:27 Nga:28]

	delete(friends, "Dung")
	fmt.Println(friends) //map[Nga:28]
	fmt.Println(people)  //map[An:30 Dung:27 Nga:28]

	delete(friends, "kirin") //not error
}
