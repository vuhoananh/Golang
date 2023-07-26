package main

import "fmt"

func main() {
	// if else
	number := 10

	// if condition { code }
	if number == 10 {
		fmt.Println("Number = 10")
	}

	// if condition { code } else { code }
	if number == 10 {
		fmt.Println("Number = 10")
	} else {
		fmt.Println("Number != 10")
	}

	if a := 100; a > 100 {
		fmt.Println("a > 100")
	} else {
		fmt.Println("a = 100")
	}
	fmt.Println("----------")

	// Switch - case
	// if else if else if else ...
	number1 := 10
	switch number1 {
	case 1, 2, 3, 10:
		fmt.Println("Number = 10")
	case 4:
		fmt.Println("Number = 10")
	case 5:
		fmt.Println("Number = 10")
	case 6:
		fmt.Println("Number = 10")
	default:
		fmt.Println("Unknown")
	}

	number2 := 10
	switch {
	case number2 > 10:
		fmt.Println("Number > 10")
	case number2 == 10:
		fmt.Println("Number = 10")
	}
	fmt.Println("----------")

	// Fallthrough, break, goto
	number3 := 10
	switch number3 {
	case 10:
		if number3 == 10 {
			fmt.Println("Break here!")
			break
		}
		fmt.Println("Number = 10")
		fallthrough
	case 1:
		fmt.Println("Number = 1")
		fallthrough
	case 2:
		fmt.Println("Number = 2")
		fallthrough
	case 3:
		fmt.Println("Number = 3")
	}

	number4 := 10
	switch number4 {
	case 10:
		if number4 == 10 {
			goto handleNumberEqual10
		}
	handleNumberEqual10:
		fmt.Println("Handle for case = 10")
	case 1:
		fmt.Println("Number = 1")
		fallthrough
	case 2:
		fmt.Println("Number = 2")
		fallthrough
	case 3:
		fmt.Println("Number = 3")
	}
	fmt.Println("----------")

	// Loops
	// for init; condition; post
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("----------")

	// break, continue
	for i := 0; i < 10; i++ {
		if i == 4 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("----------")

	for i := 0; i < 10; i++ {
		if i == 4 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("----------")

	// while
	j := 0
	for j < 10 {
		fmt.Println(j)
		j++
	}
	fmt.Println("----------")

	// Infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }

	for i, j := 1, 2; i < 10 && j < 10; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}
}
