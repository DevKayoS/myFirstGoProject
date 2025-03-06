package pacote

import "fmt"

func array() {
	const lengthArr = 10

	arr := [lengthArr]int{1, 2, 3}
	arr1 := [10]int{5: 399, 7: 321}

	arrString := [10]string{2: "sim sim "}

	fmt.Println(arr)
	fmt.Println(arr1)

	fmt.Println(arrString)
}
