package main

import "fmt"

func main() {
	// title1, author1, year1 := "The Divine Comedy", "Date Aligheri", 1320
	// title2, author2, year2 := "Macbeth", "William Shakespear", 1606

	// fmt.Println("Book1: ", title1, author1, year1)
	// fmt.Println("Book2: ", title2, author2, year2)

	type book struct {
		title  string
		author string
		year   int
	}

	// type book1 struct {
	// 	titles, author string
	// 	year           int
	// }

	// myBook1 := book{"The Divine Comedy", "Date Aligheri", 1320}
	// fmt.Println(myBook1)

	// bestBook := book{title: "Hello", author: "Kirrin", year: 1997}

	// fmt.Println(bestBook) //{Hello Kirrin 1997}

	// aBook := book{title: "Halo"}

	// fmt.Printf("%#v\n", aBook)//main.book{title:"Halo", author:"", year:0}

	lastBook := book{title: "last Bookk"}
	// fmt.Println(lastBook.title) //last Bookk

	// page := lastBook.pages//lastBook.pages undefined (type book has no field or method pages)

	lastBook.author = "Dung"
	lastBook.year = 1998
	fmt.Printf("%+v\n", lastBook) //{title:last Bookk author:Dung year:1998}

	//Compare struct
	aBook := book{title: "aBook", author: "Kirrin", year: 1997}

	fmt.Println(aBook == lastBook) //false

	myBook := aBook
	myBook.year = 1235
	fmt.Println(myBook, aBook) //{aBook Kirrin 1235} {aBook Kirrin 1997}
}
