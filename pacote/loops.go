package pacote

import (
	"fmt"
	"sync"
)

func loops() {
	arr := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, value := range arr {
		fmt.Println(value)
	}

	const n = 10

	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		// i := i
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}
