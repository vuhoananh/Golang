package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func createPizza(pizza int) {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Printf("Created Pizza %d \n", pizza)
}

func timeTrack(start time.Time, fuctionName string) {
	elapsed := time.Since(start)
	fmt.Println(fuctionName, "took", elapsed)
}

func main() {
	defer timeTrack(time.Now(), "Build Pizzas")

	number_of_ovens := 3
	runtime.GOMAXPROCS(number_of_ovens)
	wg.Add(number_of_ovens)

	for i := 0; i < number_of_ovens; i++ {
		go createPizza(i)
	}
	wg.Wait()
}
