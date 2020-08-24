package main

import "fmt"

func main() {
	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",

		age: 39,
	}
	fmt.Printf("%#v\n", diana)                   //struct { firstName string; lastName string; age int }{firstName:"Diana", lastName:"Muller", age:39}
	fmt.Printf("Diana 's age: %d \n", diana.age) //Diana 's age: 39

	type Book struct {
		string
		float64
		bool
	} //anonymous struct
	b1 := Book{"The magic", 15.43, true}
	fmt.Println(b1)       //{The magic 15.43 true}
	fmt.Printf("%#v", b1) //main.Book{string:"The magic", float64:15.43, bool:true}

	fmt.Println(b1.string) //The magic

	type Employee struct {
		name   string
		salary float64
		bool
	}
	e := Employee{"John", 19230.0, true}
	fmt.Printf("%#v", e) //main.Employee{name:"John", salary:19230, bool:true}
	e.bool = false
	fmt.Printf("%#v", e) //main.Employee{name:"John", salary:19230, bool:false}

}
