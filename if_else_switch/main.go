package main

import (
	"fmt"
	"time"
)

func main() {
	//switch
	// language := "java"
	// switch language {
	// case ".Net":
	// 	fmt.Println("You are learning .Net")
	// 	break

	// case "Go", "golang":
	// 	fmt.Println("You are learning Golang")
	// 	break

	// default:
	// 	fmt.Println("You are learning other language")
	// 	break
	// }

	// n := 5
	// switch true {
	// case n%2 == 0:
	// 	fmt.Println("Even number")
	// 	break
	// case n%2 != 0:
	// 	fmt.Println("Odd number")
	// 	break
	// }

	hour := time.Now().Hour()
	fmt.Println(hour)

	switch {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evenning")
	}

}
