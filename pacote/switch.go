package pacote

import (
	"fmt"
	"math"
	"time"
)

func main() {
	do(1)
	secondDo(2)
	fmt.Println(isWeekend(time.Now()))
}

func do(x int) {
	switch x {
	case 1:
		fmt.Println("nice dms", 1)
	case 2:
		fmt.Println("segunda parte", 2)
	case 3:
		fmt.Println("terceira parte", 3)
	default:
		fmt.Println("tem nada nao")
	}
}

func secondDo(x int) {
	switch {
	case x == 1:
		fmt.Println("danado")
	case x < 0:
		fmt.Println("numero negativo")
	case x > 1:
		fmt.Println("numero positivo superior a 1")
	}
}

func isWeekend(x time.Time) bool {
	switch {
	case x.Weekday() > 0 && x.Weekday() < 6:
		fmt.Println("como que retorna essa parada", x)
		return false
	default:
		fmt.Println("fim de semana ja chegou danado")
		return true
	}
}

func extras() {
	switch x := math.Sqrt(4); x {
	case 2:
		fmt.Println("resultado dessa raiz e 2")
	default:
		fmt.Println("num compensa dog")
	}
}

func isWeekendPt2(x time.Time) bool {
	switch x.Weekday() {
	case time.Sunday, time.Saturday:
		return true
	default:
		return false
	}
}
