package main

import (
	"fmt"
)

func main() {
	// s1 := []int{10, 20, 30, 40, 50} //[10 20 30 40 50]
	// s3, s4 := s1[0:2], s1[1:3]      //[10 20 ], [20 30]

	// s3[1] = 600

	// fmt.Println(s1) //[10 600 30 40 50]
	// fmt.Println(s4) //[600 30]]

	// arr1 := [4]int{10, 20, 30, 40}
	// slice1, slice2 := arr1[0:2], arr1[1:3]

	// arr1[1] = 2
	// fmt.Println(arr1, slice1, slice2)//[10 2 30 40] [10 2] [2 30]]

	// cars := []string{"Ford", "Honda", "BMV", "Audi"}
	// newCars := []string{}
	// newCars = append(newCars, cars[0:2]...)
	// cars[0] = "lambo"
	// fmt.Println(cars, newCars)//[lambo Honda BMV Audi] [Ford Honda]

	// s1 := []int{10, 20, 30, 40, 50}
	// newSlice := s1[0:3]
	// fmt.Println(len(newSlice), cap(newSlice)) //3 5   : 5 is len of backing arr
	// newSlice = s1[2:5]
	// fmt.Println(len(newSlice), cap(newSlice)) //3 5   : 5 is len of backing arr
	// fmt.Printf("%p \n", &s1)                  //0xc000004480 :adrress of slice s1

	// fmt.Printf("%p  %p \n", &s1, &newSlice) //0xc000004480  0xc0000044a0

	// newSlice[0] = 1000
	// fmt.Println("s1:", s1) //s1: [10 20 1000 40 50]

	// a := [5]int{1, 2, 3, 4, 5}
	// s := []int{1, 2, 3, 4, 5}
	// fmt.Printf("array's size in bytes: %d \n", unsafe.Sizeof(a))//array's size in bytes: 40
	// fmt.Printf("slice's size in bytes: %d \n", unsafe.Sizeof(s))//slice's size in bytes: 24

	var nums []int
	fmt.Printf("%#v\n", nums)                                      //[]int(nil)
	fmt.Printf("length: %d, Capacity: %d\n", len(nums), cap(nums)) //length: 0, Capacity: 0

	nums = append(nums, 1, 2)
	fmt.Printf("length: %d, Capacity: %d\n", len(nums), cap(nums)) //length: 2, Capacity: 2

	nums = append(nums, 3)
	fmt.Printf("length: %d, Capacity: %d\n", len(nums), cap(nums)) //length: 3, Capacity: 4

	nums = append(nums, 4)
	fmt.Printf("length: %d, Capacity: %d\n", len(nums), cap(nums)) //length: 4, Capacity: 4

	nums = append(nums, 100)
	fmt.Printf("length: %d, Capacity: %d\n", len(nums), cap(nums)) //length: 5, Capacity: 8 //Capacity=oldCapacity *2

	nums = append(nums[0:4], 500, 600, 700, 800, 900)
	fmt.Printf("length: %d, Capacity: %d\n", len(nums), cap(nums)) //length: 9, Capacity: 16 //Capacity=oldCapacity *2

	letters := []string{"A", "B", "C", "D", "E", "F"}

	letters = append(letters[:5], "X", "Y")
	fmt.Println(letters, len(letters), cap(letters)) //[A B C D E X Y] 7 12

}
