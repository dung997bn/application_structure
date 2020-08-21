package main

import (
	"fmt"
	"strings"
)

func main() {
	// 	s1 := "Hello Kirin"
	// 	fmt.Println(s1)

	// 	fmt.Println("He says \"Ok\"") //He says "Ok"
	// 	fmt.Println(`He says "Ok"`)   //He says "Ok" this call raw string

	// 	s2 := "Price: 1000 \nBrand: Nike"
	// 	fmt.Println(s2)
	// 	fmt.Println(`
	// Price: 1000
	// Brand: Nike`)

	// 	fmt.Println(`C:\User\Document`)   //C:\User\Document
	// 	fmt.Println("C:\\User\\Document") //C:\User\Document

	// var s3 = "I love " + "Go " + "Programming"
	// fmt.Println(s3 + "!")   //I love Go Programming!
	// fmt.Printf("%s \n", s3) //I love Go Programming
	// fmt.Printf("%q \n", s3) //"I love Go Programming"

	// var1, var2 := 'a', 'b'
	// fmt.Printf("Type: %T, Value: %d \n", var1, var2)

	// str := "âtă"
	// fmt.Println(len(str))
	// fmt.Println("Byte (not rune) at position 1:", str[1]) // 162

	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("%c", str[i]) //Ã¢tÄ
	// }

	// fmt.Println("\n" + strings.Repeat("#", 20)) //####################
	// for i := 0; i < len(str); {
	// 	r, size := utf8.DecodeRuneInString(str[i:])
	// 	fmt.Printf("%c", r) //âtă
	// 	i += size
	// }

	// fmt.Println("\n" + strings.Repeat("#", 20)) //####################

	// for _, r := range str {
	// 	fmt.Printf("%c", r)//âtă
	// }

	// s1 := "Golang"
	// fmt.Println(len(s1)) //6

	// name := "tajị"
	// n := utf8.RuneCountInString(name)
	// m := utf8.RuneCountInString("@")
	// fmt.Println(len(name), len("@"))//6,1
	// fmt.Println(n, m)//4,1

	// s1 := "I Love Golang"
	// fmt.Println(s1[2:5]) //Lov

	// rs1 := []rune(s1)
	// fmt.Println(string(rs1))//I Love Golang

	// s2 := "かんたん選挙会計"
	// fmt.Println(s2[0:2]) //�
	// fmt.Println(s2[0:3]) //か

	// rs := []rune(s2)
	// fmt.Printf("%T\n", rs) //[]int32

	// fmt.Println(string(rs[0:3])) //かんた

	// fmt.Println(string(rs)) //かんたん選挙会計

	//Contains
	p := fmt.Println
	// result := strings.Contains("I love Go Programming", "Love") //false
	// //result := strings.Contains("I love Go Programming", "love") //true
	// p(result)

	//ContainsAny
	//result := strings.ContainsAny("I love Go Programming", "wz")//false
	// result := strings.ContainsAny("I love Go Programming", "vo") //true
	// p(result)

	//ContainsRune
	//result := strings.ContainsRune("I love go Programming", 'G') //false
	// result := strings.ContainsRune("I love go Programming", 'g') //true
	// p(result)

	// n := strings.Count("cheese", "e")//3
	// n := strings.Count("cheese", "") //7
	// p(n)

	// p(strings.ToLower("Go Jave .Net")) //go jave .net
	// p(strings.ToUpper("go jave .net")) //GO JAVE .NET

	// p("go" == "go") //true
	// p("Go" == "go") //false

	// p(strings.ToLower("Go") == "go")  //true
	// p(strings.ToLower("Go") == "go ") //false

	// myStr := strings.Repeat("ab", 10)//abababababababababab
	// p(myStr)

	//myStr := strings.Replace("192.168.12.34", ".", ":", 2)//192:168:12.34
	// myStr := strings.Replace("192.168.12.34", ".", ":", -2) //192:168:12:34 //replace all if n<0
	// p(myStr)

	myStr := strings.ReplaceAll("192.168.12.34", ".", ",") //192,168,12,34   //replace all
	p(myStr)

	s := strings.Split("a,b,c", ",")
	fmt.Printf("%T\n", s)  //[]string
	fmt.Printf("%#v\n", s) //[]string{"a", "b", "c"}

	z := strings.Split("Go for Go!", "")
	fmt.Printf("%#v\n", z) //[]string{"G", "o", " ", "f", "o", "r", " ", "G", "o", "!"}

	x := []string{"I", "Learn", "Golang"}
	//myStrJoin := strings.Join(x, "-")//I-Learn-Golang
	myStrJoin := strings.Join(x, "aaaa") //IaaaaLearnaaaaGolang
	p(myStrJoin)

	w := "orange Green Blue yellow"
	fields := strings.Fields(w)

	fmt.Printf("%T\n", fields)  //[]string
	fmt.Printf("%#v\n", fields) //[]string{"orange", "Green", "Blue", "yellow"}

	x1 := strings.TrimSpace("\t Goodbye Windows, Linux, Welcom Mac \n")
	fmt.Printf("%q\n", x1) //"Goodbye Windows, Linux, Welcom Mac"

	x2 := strings.Trim("....halo Kirin!!!!?", ".!?")
	fmt.Printf("%q\n", x2) //"halo Kirin"

}
