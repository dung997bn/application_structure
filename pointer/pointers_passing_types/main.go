package main

import "fmt"

func changeValues(quality int, price float64, name string, sold bool) {
	quality = 3
	price = 500.4
	name = "Mobile Phone"
	sold = false
}

func changeValuesByPointer(quality *int, price *float64, name *string, sold *bool) {
	*quality = 3
	*price = 500.4
	*name = "Mobile Phone"
	*sold = false
}

type product struct {
	price       float64
	productName string
}

func changeProduct(p product) {
	p.price = 30
	p.productName = "Bikecycle"
}

func changeProductByPointer(p *product) {
	(*p).price = 300 //equal p,price=300
	(*p).productName = "Car"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

func main() {
	// quality, price, name, sold := 5, 300.4, "Laptop", true
	// fmt.Println("Before calling changesValue():", quality, price, name, sold) //Before calling changesValue(): 5 300.4 Laptop true
	// changeValues(quality, price, name, sold)
	// fmt.Println("After calling changesValue():", quality, price, name, sold) //After calling changesValue(): 5 300.4 Laptop true

	// changeValuesByPointer(&quality, &price, &name, &sold)

	// fmt.Println("After calling changeValuesByPointer():", quality, price, name, sold) //After calling changeValuesByPointer(): 3 500.4 Mobile Phone false

	// gift := product{
	// 	price:       100,
	// 	productName: "Watch",
	// }
	// changeProduct(gift)
	// fmt.Println(gift) //{100 Watch}

	// fmt.Println("Before calling changeProductByPointer():", gift) //Before calling changeProductByPointer(): {100 Watch}
	// changeProductByPointer(&gift)
	// fmt.Println("After calling changeProductByPointer():", gift) //After calling changeProductByPointer(): {300 Car}

	prices := []int{1, 2, 3}

	fmt.Println("Before calling changeSlice():", prices) //Before calling changeSlice(): [1 2 3]
	changeSlice(prices)
	fmt.Println("After calling changeSlice():", prices) //After calling changeSlice(): [2 3 4]

	myMap := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println("Before calling changeMap():", myMap) //Before calling changeMap(): map[a:1 b:2 c:3]
	changeMap(myMap)
	fmt.Println("After calling changeMap():", myMap) //After calling changeMap(): map[a:10 b:20 c:30]

	//Function change value of int, string, bool, float, struct ,..etc best by pointer
	//Slices and map are modified poiter, so don't need pass pointer of slice and map to change value

}
