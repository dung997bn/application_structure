package main

import "fmt"

func main() {
	type Contact struct {
		email, address string
		phone          int
	}
	type Employee struct {
		name        string
		salary      float64
		contactInfo Contact
	}

	john := Employee{
		name: "John", salary: 192323.43, contactInfo: Contact{
			email:   "john@gmail.com",
			address: "Londo",
			phone:   948364834834,
		},
	}
	fmt.Printf("%#v \n", john) //main.Employee{name:"John", salary:192323.43, contactInfo:main.Contact{email:"john@gmail.com", address:"Londo", phone:948364834834}}

	fmt.Println("Employee 's Email:", john.contactInfo.email) //Employee 's Email: john@gmail.com

	john.contactInfo.email = "newemail@gmail.com"

	fmt.Println("Employee 's Email:", john.contactInfo.email) //Employee 's Email: newemail@gmail.com

	myContact := Contact{"dung@gmail.com", "Ha noi", 99302324}
	fmt.Println(myContact) //{dung@gmail.com Ha noi 99302324}
}
