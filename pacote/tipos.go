package pacote

import "fmt"

//bool

// int int8 int16 int32 int64
// uint uint8 uint16 uint32 uint64 uintptr

// byte
// rune

// float32 float64

// complex64 complex128

// string

func tipos() {
	const x = 10
	takeInt32(12)
	takeInt64(x)
	takeString("dsdasdas")
	takeFloat64(x)
}

func takeInt32(x int32) {
	fmt.Println(x)
}

func takeInt64(x int64) {
	fmt.Println(x)
}

func takeString(x string) {
	fmt.Println(x)
}

func takeFloat64(x float64) {
	fmt.Println(x)
}
