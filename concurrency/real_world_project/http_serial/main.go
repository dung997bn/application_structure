package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func checkAndSaveBody(url string) {
	rest, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN\n", url)
	} else {
		defer rest.Body.Close()
		fmt.Printf("%s -> Status Code: %d \n", url, rest.StatusCode)
		if rest.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(rest.Body)
			file := strings.Split(url, "//")[1] //http://www.google.com
			file += ".txt"
			fmt.Printf("Writing response body to %s\n", file)

			err = ioutil.WriteFile(file, bodyBytes, 0064)

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
func main() {
	urls := []string{"https://golang1234.org", "https://www.google.com/a.html", "https://www.medium.com"}
	for _, url := range urls {
		checkAndSaveBody(url)
		fmt.Println(strings.Repeat("#", 20))
	}
}
