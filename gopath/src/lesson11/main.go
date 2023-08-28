package main

import "fmt"

func main() {
	var myMap = make(map[string]int)
	fmt.Println(myMap)

	if myMap == nil {
		fmt.Println("myMap = nil")
	} else {
		fmt.Println("myMap != nil")
	}

	fmt.Println()

	var myMap1 map[string]int
	fmt.Println(myMap1)

	if myMap1 == nil {
		fmt.Println("myMap1 = nil")
	}

	fmt.Println()

	// Khai bao voi gia tri khoi tao
	myMap2 := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}
	fmt.Println(myMap2)

	// Them 1 phan tu vao map
	myMap2["key4"] = 4
	myMap2["key5"] = 5
	fmt.Println(myMap2)

	// delete 1 phan tu trong map = delete (map, key)
	delete(myMap2, "key1")
	fmt.Println(myMap2)

	// lenth cua map
	fmt.Println(len(myMap2))

	// map la reference type
	myMap3 := myMap2
	fmt.Println(myMap3)

	myMap2["key5"] = 1000
	fmt.Println(myMap3)

	// Truy cap 1 phan tu trong map
	value := myMap3["key5"]
	fmt.Println(value)

	value1, found := myMap3["key1000"]
	if found {
		fmt.Println(value1)
	} else {
		fmt.Println("Ko co gia tri voi key = 1000")
	}

	// Trong map ko co toan tu ==
	// if myMap2 == myMap3 {
	// 	fmt.Println(myMap2)
	// }
}
