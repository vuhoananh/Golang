package main

import "fmt"

// func func_name(params) return_type { code }
func chao() {
	fmt.Println("Hello")
}

func Chao(name string) {
	fmt.Println("Hello", name)
}

func greeting(name string) string {
	result := fmt.Sprintf("Hello %s", name)
	return result
}

// Multiple return value
func rectInfo(w, h int) (int, int, int) {
	area := w * h
	return w, h, area
}

// Named return value
func rectInfo1(w, h int) (width int, height int, isSquare bool) {
	isSquare = w == h
	return w, h, isSquare
}

func main() {
	chao()

	Chao("Ryan")

	result := greeting("Alice")
	fmt.Println(result)

	w, h, area := rectInfo(100, 200)
	fmt.Println("width=", w)
	fmt.Println("height=", h)
	fmt.Println("area=", area)

	w, h, isSquare := rectInfo1(200, 100)
	if isSquare {
		fmt.Println("This is square")
	} else {
		fmt.Println("width=", w)
		fmt.Println("height=", h)
	}

}
