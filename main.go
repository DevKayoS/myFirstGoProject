package main

import (
	"fmt"
	"sync"
)

func main() {
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
