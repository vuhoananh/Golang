package main

import "fmt"

// Interface
type Animal interface {
	speak()
}

// Multiple interfaces
type Movement interface {
	move()
}

// Embeb interface
type NextAnimal interface {
	Animal
	Movement
}

type Dog struct{}

func (d Dog) speak() {
	fmt.Println("woawww woawww")
}

func (d Dog) move() {
	fmt.Println("Dog chay bang 4 chan")
}

// Empty interface
func goOut(i interface{}) {
	fmt.Println(i)
}

type Data struct{
	index int
}

func main() {
	var animal Animal

	animal = Dog{}

	animal.speak()

	fmt.Println("---------------")

	dog := Dog{}

	var a Animal = dog
	a.speak()

	var m Movement = dog
	m.move()

	fmt.Println("---------------")

	var na NextAnimal = dog
	na.speak()
	na.move()

	fmt.Println("---------------")

	goOut(10)
	goOut(10.12)

	fmt.Println("---------------")

	d:=Data {100}
	goOut(d)
}
