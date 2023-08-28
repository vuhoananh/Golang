package main

import (
	"fmt"
)

// Channel

func main() {
	// Unbufferred channel

	ch := make(chan int)

	go func() {
		ch <- 100
		fmt.Println("Sent")
	}()

	fmt.Println(<-ch)
	fmt.Println("Done")

	fmt.Println()

	// Bufferred channel

	ch1 := make(chan int, 2)

	ch1 <- 1
	ch1 <- 2

	close(ch1)

	fmt.Println(<-ch1)
	fmt.Println(<-ch1)
	fmt.Println(<-ch1)

	fmt.Println()

	// Close channel
	queue := make(chan int, 10)

	go func() {
		for i := 0; i < 10; i++ {
			if i > 5 {
				close(queue)
				break
			} else {
				queue <- i
			}
		}
	}()

	for value := range queue {
		fmt.Println(value)
	}

	fmt.Println()

	// Select channel
	queue2 := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			queue2 <- i
		}

		done <- true
	}()

	for {
		select {
		case v := <-queue2:
			fmt.Println(v)

		case <-done:
			fmt.Println("Done")
			return
		}
	}
}
