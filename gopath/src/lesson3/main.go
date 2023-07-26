package main

import (
	"fmt"
	"math"
	"math/bits"
)

func main() {
	// Bool
	var myBool bool = true
	fmt.Println(myBool)

	var mySecondBool bool = false
	fmt.Println(mySecondBool)

	// String
	var myString string = "Hello"
	fmt.Println(myString)

	// Int
	var myInt int = 12346
	fmt.Println(myInt)

	fmt.Println("----------")

	// Int 8, 16, 32, 64
	// 1. Range
	// Int8
	fmt.Println(math.MinInt8)
	fmt.Println(math.MaxInt8)

	// Int16
	fmt.Println(math.MinInt16)
	fmt.Println(math.MaxInt16)

	// Int32
	fmt.Println(math.MinInt32)
	fmt.Println(math.MaxInt32)

	// Int64
	fmt.Println(math.MinInt64)
	fmt.Println(math.MaxInt64)

	fmt.Println("----------")

	//2. Bits
	fmt.Println(bits.OnesCount8(math.MaxUint8))
	fmt.Println(bits.OnesCount16(math.MaxUint16))
	fmt.Println(bits.OnesCount32(math.MaxUint32))
	fmt.Println(bits.OnesCount64(math.MaxUint64))

	fmt.Println("----------")

	// Uint
	var myUint uint = 10
	fmt.Println(myUint)

	fmt.Println("----------")

	// Byte
	var myByte byte = 255
	fmt.Println(myByte)
	fmt.Printf("%T", myByte)
	// fmt.Println(math.MaxUint8)

	fmt.Println()
	fmt.Println("----------")

	var a byte = 'A'
	fmt.Println(a)
	fmt.Printf("%X", a)
	fmt.Println()

	fmt.Println("----------")

	// Float
	var myFloat float64 = 10.01
	fmt.Println(myFloat)

	fmt.Println("----------")

	// Complex z = a + bi
	var z1 = 10 + 1i
	fmt.Println(z1)

	var z2 = 10 + 1i
	fmt.Println(z1 + z2)

	fmt.Println("----------")

	// Rune
	var myRune string = "Nhẫn"
	runes := []rune(myRune)
	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c", runes[i])
	}

	fmt.Println()

	var myRune2 rune = 'A'
	fmt.Printf("%T", myRune2)

	fmt.Println()

	var myRune3 rune = '🅐'
	fmt.Printf("%c - %U", myRune3, myRune3)
	fmt.Println()

	// Type conversions
	var b int = 1
	var c float64 = 2.1
	fmt.Println(float64(b) + c)

	// Constans
	const PI = 3.14159
	// PI := 3.14159
	fmt.Println(PI)

	// Uintptr
	// key - value
	// key = uintptr
	// value = object (*point)

	// key = (*point) -> uintptr
	// map.put(key, value)
}
