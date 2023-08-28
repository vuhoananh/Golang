package main

import "fmt"

// Struct

type Student struct {
	id   int
	name string
}

type StudentInfo struct {
	class string
	email string
	age   int
}

type Students struct {
	id   int
	name string
	info StudentInfo
}

func main() {
	st1 := Student{
		id:   100,
		name: "Ryan",
	}

	fmt.Println(st1)
	fmt.Println(st1.id)
	fmt.Println(st1.name)

	fmt.Println()

	st2 := Student{200, "Robin"}
	fmt.Println(st2)
	fmt.Println(st2.id)
	fmt.Println(st2.name)

	fmt.Println()

	var st3 Student = struct {
		id   int
		name string
	}{
		id:   300,
		name: "Bill",
	}
	fmt.Println(st3)

	fmt.Println()

	// anonymous struct
	var anonymous = struct {
		email string
		age   int
	}{
		email: "ryan@gmail.com",
		age:   30,
	}
	fmt.Println(anonymous)

	// pointer tro toi struct
	pointer := &Student{
		id:   999,
		name: "Robin",
	}
	fmt.Println(&pointer)
	fmt.Println(pointer.id)
	fmt.Println(pointer.name)

	fmt.Println()

	// anonymous fields
	type Noname struct {
		string
		int
	}

	n := Noname{"Nguyen Van A", 6666}
	fmt.Println(n)

	fmt.Println()

	//struct long struct - Nested struct
	student := Students{
		id:   123,
		name: "Ryan",
		info: StudentInfo{
			class: "ql091",
			email: "ryan@gmail.com",
			age:   30,
		},
	}
	fmt.Println(student)

	fmt.Println()

	// Compare 2 struct
	type struct1 struct {
		id   int
		name string
	}

	s1 := struct1{1, "A"}
	s2 := struct1{1, "A"}

	if s1 == s2 {
		fmt.Println("s1 == s2")
	} else {
		fmt.Println("s1 != s2")
	}

	fmt.Println()

	// zero value
	var student2 Student

	student2.name = "Ryan"

	fmt.Println(student2)
}
