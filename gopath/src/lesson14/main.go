package main

import "fmt"

type Student struct {
	name string
}

// Define Method

// func (receiver type) methodName() return { // code }

func (s Student) getName() string {
	return s.name
}

// 1. Value Receiver
func (s Student) changeName() {
	fmt.Printf("p2 = %p", &s)
	s.name = "Robin"
	fmt.Println()
}

// 2. Pointer Receiver
func (s *Student) changeName2() {
	fmt.Printf("p4 = %p", s)
	s.name = "Robin"
	fmt.Println()
}

// non-struct
type String string

func (s String) append(str string) string {
	return fmt.Sprintf("%s%s",s, str)
}

func main() {
	student := Student{"Ryan"}
	fmt.Printf("p1 = %p", &student)
	fmt.Println()

	student.changeName()
	fmt.Println(student.name)

	fmt.Println("---------------")

	student2 := &Student{"Ryan"}
	fmt.Printf("p3 = %p", &student2)
	fmt.Println()

	student2.changeName2()
	fmt.Println(student2.name)

	s1 := String("a")
	newStr:=s1.append("b")

	fmt.Println(newStr)
}
