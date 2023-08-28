package main

import "fmt"

func main() {
	a := 100.1

	var pointer *float64

	pointer = &a
	fmt.Println(pointer)
	fmt.Printf("%T", pointer)
	fmt.Println()

	fmt.Println("---------------")

	b := 123

	pointer2 := new(int) //<=> var pp2 *int

	pointer2 = &b

	fmt.Println(pointer2)
	fmt.Printf("%T", pointer2)
	fmt.Println()

	fmt.Println("---------------")

	//zero value
	var pointer3 *int

	pointer4 := new(int)

	fmt.Println(pointer3)
	fmt.Println(pointer4)

	fmt.Println("---------------")

	c := 666

	var pointer5 *int

	pointer5 = &c

	fmt.Println(c)

	*pointer5 = 999
	fmt.Println(c)

	fmt.Println("---------------")

	// demo pointer -> array
	array := [3]int {1,2,3}

	var pointer6 *[3]int

	pointer6 = &array

	fmt.Println(pointer6)
	fmt.Printf("%T",pointer6)
	fmt.Println()

	fmt.Println("---------------")

	d := 333

	var pointer7 *int = &d

	apllyPointer(pointer7)

	fmt.Println(d)
}

func apllyPointer (pointer *int) {
	*pointer = 777
}