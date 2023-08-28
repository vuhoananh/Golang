package main

import "fmt"

// Array
func main() {
	// Khai bao array
	var myArray [4]int
	fmt.Println(myArray)

	// gan gia tri
	myArray[0] = 123
	fmt.Println(myArray)

	myArray[3] = 456
	fmt.Println(myArray)

	// Khai bao array co gia tri khoi tao
	array := [3]int{1, 2, 3}
	fmt.Println(array)

	array2 := [3]int{100}
	fmt.Println(array2)

	//size mang
	fmt.Println(len(myArray))

	// Khai bao array ko can set size
	array3 := [...]string{"Vinfast", "Honda", "Huyndai", "Tesla", "1", "2", "3"}
	fmt.Println(array3)
	fmt.Println(array3[0])

	// Array la value type ko phai la reference type
	countries := [...]string{"VN", "US", "PR"}
	copyCountries := countries

	copyCountries[0] = "CN"

	fmt.Println(countries)
	fmt.Println(copyCountries)

	// Loop array
	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])

	}

	for index, value := range countries {
		fmt.Printf("i=%d value=%s", index, value)
		fmt.Println()
	}

	// blank identifier
	for _, value := range countries {
		fmt.Printf("value=%s", value)
		fmt.Println()
	}

	// Array 2 chieu
	matrix := [4][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
		{7, 8},
	}
	fmt.Println(matrix)

	for i := 0; i < 4; i++ {
		for j := 0; j < 2; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
