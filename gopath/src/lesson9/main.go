package main

import "fmt"

func main() {
	// Khai bao slice
	var slice []int
	fmt.Println(slice)

	// Khai bao slice va khoi tao slice
	var slice1 = []int{1, 2, 3, 4}
	fmt.Println(slice1)

	// Tao 1 slice tu 1 array
	var array = [4]int{1, 2, 3, 4}
	slice2 := array[1:3]
	fmt.Println(slice2)

	slice3 := array[:]
	fmt.Println(slice3)

	slice4 := array[2:]
	fmt.Println(slice4)

	slice5 := array[:3]
	fmt.Println(slice5)

	// Tao 1 slice tu 1 slice khac
	var slice6 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice7 := slice6
	fmt.Println(slice7)

	slice8 := slice6[1:]
	fmt.Println(slice8)

	// Slice => referance type
	var array1 = []int{1, 2, 3, 4, 5}

	slice9 := array1[:]
	slice9[0] = 999

	fmt.Println(array1)
	fmt.Println(slice9)

	// Length vs capacity cua slice
	// len: so luong phan tu cua slice
	// cap: so luong phan tu cua underlying array bat dau tu vi tri start khi ma slice dc tao
	countries := [...]string{"VN", "USA", "CANADA", "FRANCE", "CHINA"}

	slice10 := countries[2:3]
	fmt.Println(slice10)

	fmt.Println(len(slice10))
	fmt.Println(cap(slice10))

	// make
	slice11 := make([]int, 2, 5)
	fmt.Println(slice11)
	fmt.Println(len(slice11))
	fmt.Println(cap(slice11))

	slice12 := make([]int, 2)
	fmt.Println(slice12)
	fmt.Println(len(slice12))
	fmt.Println(cap(slice12))

	// append
	var slice13 []int
	slice13 = append(slice13, 100)
	fmt.Println(slice13)

	// copy
	src := []string{"A", "B", "C", "D"}

	dest := make([]string, 2)
	number := copy(dest, src)

	fmt.Println(number)
	fmt.Println(dest)

	//delete
	src = append(src[:1], src[2:]...)
	fmt.Println(src)
}
