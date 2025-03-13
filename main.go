package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Jogo de adivinhacao")
	fmt.Println("Um Numero aleatorio sera sorteado. Tente acertar. O numero e um inteiro entre 0 e 100")

	x := rand.Int64N(101)
	//lendo o terminal
	scanner := bufio.NewScanner(os.Stdin)
	chutes := [10]int64{}

	for i := range chutes {
		fmt.Println("Qual e o seu chute?")
		// chamando a funcao e pegando os dados do terminal
		scanner.Scan()
		chute := scanner.Text() //pega o valor do terminal
		chute = strings.TrimSpace(chute)

		chuteInt, err := strconv.ParseInt(chute, 10, 64)
		if err != nil {
			fmt.Println("O seu chute tem que ser um numero inteiro")
			return
		}
		switch {
		case chuteInt < x:
			fmt.Println("Voce errou. O numero sorteado e maior que: ", chuteInt)
		case chuteInt > x:
			fmt.Println("Voce errou. O numero sorteado e menor que: ", chuteInt)
		case chuteInt == x:
			fmt.Printf(
				"Parabens!, voce  acertou o numero! o numero era: %d\n"+
					"Voce teve %d tentativas.\n"+
					"Essas foram as suas tentativas: %v\n",
				x, i+1, chutes[:i],
			)

			return
		}

		chutes[i] = chuteInt
	}

	fmt.Printf(
		"Infelizmente, voce nao acertou o numero, que era: %d\n"+
			"Voce teve 10 tentativas.\n"+
			"Essas foram as suas tentativas: %v\n",
		x, chutes,
	)

}
