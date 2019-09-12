package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {

	var sliceString []int

	//O conhecido

	for i := 0; i < 100; i++ {
		sliceString = append(sliceString, rand.Intn(100))
	}

	//O por chave

	for indice, valor := range sliceString {
		fmt.Println("Indice: " + strconv.Itoa(indice) + " Valor: " + strconv.Itoa(valor))
	}

	//Why while ???
	// O for pode ter nenhum, 1 ou 3 argumentos :O

	contador := 0
	for {
		contador++

		if contador%2 == 0 {
			fmt.Println("Pula a mensagem!")
			continue
		}

		if contador < 100 {
			fmt.Println("Looping infinito até 100. Estou em " + strconv.Itoa(contador))

		}
		if contador >= 100 {
			break
		}
	}

	// Ou

	contador = 0
	for contador <= 100 {
		contador++

		if contador%2 == 0 {
			fmt.Println("Pula a mensagem!")
			continue
		}

		if contador < 100 {
			fmt.Println("Looping infinito até 100. Estou em " + strconv.Itoa(contador))
		}
	}

}
