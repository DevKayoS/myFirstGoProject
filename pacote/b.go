package pacote

import (
	"MyFirstGoProject/pacote/internal/foo"
	"fmt"
)

var Bar string = "hello, Bar"

func PrintMinha() {
	fmt.Println(foo.Minha)
}
