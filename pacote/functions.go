package pacote

func somar(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func swap(a, b int) (int, int) {
	return b, a
}

func dividir(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b
	return res, rem
}

// argumento variavel
func media(nums ...int) int {
	var out int

	for _, n := range nums {
		out += n
	}

	return out / len(nums)
}
