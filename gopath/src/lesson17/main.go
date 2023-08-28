package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func g1() {
	fmt.Println("g1")
}

func g2() {
	fmt.Println("g2")
	wg.Done()
}

func g3() {
	fmt.Println("g3")
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	// go func{}

	go g1()

	ng := runtime.NumGoroutine()
	fmt.Println(ng)

	time.Sleep(time.Second)

	fmt.Println()


	// Synchronnized goroutines

	fmt.Println("Begin")

	wg.Add(2)

	go g2()
	go g3()

	wg.Wait()

	fmt.Println("End")
}
