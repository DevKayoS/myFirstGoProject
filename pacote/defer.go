package pacote

import (
	"database/sql"
	"fmt"
	"os"
)

func algumNomedeFuncao() {
	x := doDefer()
	fmt.Println(x)
}

func doDefer() int {
	defer fmt.Println("world")
	fmt.Println("Hello")
	return 10
}

func deferFuncao() {
	x := 10
	defer func(y int) {
		fmt.Println(y)
	}(x)

	x = 50
	fmt.Println(x)
}

func doDeferFile() {
	file, err := os.Open("foo.txt")

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// if err := x(); err != nil {
	// 	panic(err)
	// }

	// if err := x(); err != nil {
	// 	panic(err)
	// }
}

var arquivos = []string{"foo.txt", "bar.txt", "baz.txt"}

func doDeferEscopo() {
	func() {
		for _, f := range arquivos {
			file, err := os.Open(f)
			if err != nil {
			}

			defer file.Close()
		}
	}()
}

func connectToDb() (*sql.DB, error) {
	db, err := sql.Open("", "")
	if err != nil {
	}

	defer db.Close()
	return db, nil
}
