package pacote

import (
	"errors"
	"fmt"
	"math"
)

func ifandElse() {
	if x := math.Sqrt(4); x < 1 {
		fmt.Println(x)
	} else if x < 1 {
		fmt.Println("numero positivo")
	} else {
		fmt.Println("caiu no else esse safado")
	}

	if err := doError(); err != nil {
		// um jeito bacana de fazer
	}
}

func doError() error {
	return errors.New("error")
}
