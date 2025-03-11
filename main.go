package main

import "fmt"

func main() {

}

func do(x any) {
	//descobrindo o tipo da variavel
	switch x.(type) {
	case string:
		fmt.Println("que cara das linguas")
	case int:
		fmt.Println("que cara de exatas")
	default:
		fmt.Println("admin por amor")
	}
}
