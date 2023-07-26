package main

import "fmt"

func main() {
	var number int
	number = 10
	fmt.Println(number)

	var number1 int = 11
	fmt.Println(number1)

	// Type Inference
	var email = "ryan@gmail.com"
	fmt.Println(email)

	// Khai bao nhieu bien
	// 1.Khai bao nhieu bien cung kieu du lieu
	var a, b int
	a = 10
	b = 11
	fmt.Println(a, b)

	var a1, b1 int = 10, 11
	fmt.Println(a1, b1)

	var a2, b2 = 10, 11
	fmt.Println(a2, b2)

	// 1.Khai bao nhieu bien khac kieu du lieu
	var (
		name string
		address string
		age int
	)
	name = "Robin"
	address = "Viet Nam"
	age = 25
	fmt.Println(name, address, age)

	var (
		name1 string = "Robin"
		address1 string = "Viet Nam"
		age1 int = 25
	)
	fmt.Println(name1, address1, age1)

	var name3 ,address3, age3 ="Robin","Viet Nam",25
	fmt.Println(name3, address3, age3)

	language := "Golang"
	fmt.Println(language)
}
