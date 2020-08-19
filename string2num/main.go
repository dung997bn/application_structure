package main

import (
	"fmt"
	"strconv"
)

func main() {
	// s := string(99)
	// fmt.Println(s) //c

	//s1 := string(99.2) //error
	//Num to string
	// var myStr = fmt.Sprintf("%f \n", 44.5)
	// fmt.Println(myStr)

	// var myStr1 = fmt.Sprintf("%d \n", 44)
	// fmt.Println(myStr1)

	// //String to num
	// s1 := "112.3" //string
	// var f1, err = strconv.ParseFloat(s1, 64)
	// _ = err
	// fmt.Println(f1)

	// //extra
	// i, err1 := strconv.Atoi("-32")
	// s2 := strconv.Itoa(20)
	// _ = err1
	// fmt.Printf("i is type %T and value is %v \n", i, i)
	// fmt.Printf("s2 is type %T and value is %q \n", s2, s2)

	s1 := "1124" //string
	var f1, err = strconv.ParseInt(s1, 10, 64)
	_ = err
	fmt.Println(f1)
}
