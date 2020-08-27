package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
)

func checkAndSaveBody(url string, c chan string) {
	rest, err := http.Get(url)
	if err != nil {
		//fmt.Println(err)
		s := fmt.Sprintf("%s is DOWN\n", url)
		s += fmt.Sprintf("Error: %v\n", err)
		c <- s //sending into the chanel
	} else {
		defer rest.Body.Close()
		s := fmt.Sprintf("%s -> Status Code: %d \n", url, rest.StatusCode)
		if rest.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(rest.Body)
			file := strings.Split(url, "//")[1] //http://www.google.com
			file += ".txt"
			s += fmt.Sprintf("Writing response body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0064)

			if err != nil {
				log.Fatal(err)
				s += "Error writing file\n"
				c <- s
			}
			s += fmt.Sprintf("%s is UP\n", url)
			c <- s
		}
	}

}
func main() {
	urls := []string{"https://golang.org123", "https://www.google.com/g.html", "https://www.medium.com"}
	//1.
	c := make(chan string)

	for _, url := range urls {
		go checkAndSaveBody(url, c)
		fmt.Println(strings.Repeat("#", 20))
	}

	fmt.Println("Number of goroutines:", runtime.NumGoroutine())

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-c)
	}
}
