package main

import (
	"fmt"
	"runtime"
	"time"
)

func f1() {
	fmt.Println("f1 (goroutine) excutioon started")
	for i := 0; i < 4; i++ {
		fmt.Println("f1, i=", i)
	}
	fmt.Println("f1 (goroutine) excutioon finished")
}

func f2() {
	fmt.Println("f2 (goroutine) excution started")
	for i := 0; i < 4; i++ {
		fmt.Println("f2, i=", i)
	}
	fmt.Println("f2 (goroutine) excution finished")
}

func main() {
	// fmt.Println("Main excution started")
	// fmt.Println("No. of CPUs", runtime.NumCPU())              //No. of CPUs 8
	// fmt.Println("No. of Goroutines:", runtime.NumGoroutine()) //No. of Goroutines: 1

	// fmt.Println("Os: ", runtime.GOOS)     //Os:  windows
	// fmt.Println("Arch: ", runtime.GOARCH) //Arch:  amd64

	// fmt.Println("GOMAXPROCS: ", runtime.GOMAXPROCS(0))//GOMAXPROCS:  8
	go f1()
	fmt.Println("No. of Goroutines after go f1():", runtime.NumGoroutine())

	f2()
	time.Sleep(time.Second * 5)

	fmt.Println("main excution stopped")
}
