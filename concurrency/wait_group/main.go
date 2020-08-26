package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 (goroutine) excutioon started")
	for i := 0; i < 4; i++ {
		fmt.Println("f1, i=", i)
		time.Sleep(time.Second * 2)
	}
	fmt.Println("f1 (goroutine) excutioon finished")
	wg.Done()
	//(*wg).Done() //be the same 	wg.Done()
}

func f2() {
	fmt.Println("f2 (goroutine) excution started")
	for i := 0; i < 4; i++ {
		fmt.Println("f2, i=", i)
	}
	fmt.Println("f2 (goroutine) excution finished")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go f1(&wg)
	fmt.Println("No. of Goroutines after go f1():", runtime.NumGoroutine())
	wg.Wait()
	f2()
	//time.Sleep(time.Second * 5)

	fmt.Println("main excution stopped")
}
