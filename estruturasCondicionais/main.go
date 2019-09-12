package main

import (
	"fmt"
	"os"
)

func main() {

	//Recebendo argumentos via linha de comando

	argumentosComNomePrograma := os.Args
	argumetosSemNomePrograma := os.Args[1:]

	fmt.Println(argumentosComNomePrograma)

	//Estruturas condicionais

	if len(argumetosSemNomePrograma) > 1 {
		fmt.Println(argumetosSemNomePrograma[0])
	}

	if len(argumetosSemNomePrograma) == 1 {

		if primeiroArgumento := argumetosSemNomePrograma[0]; len(primeiroArgumento) > 1 {
			fmt.Println("O único argumento informado foi: " + primeiroArgumento)
		}

	} else if len(argumetosSemNomePrograma) == 2 {
		fmt.Println("Você informou 2 argumentos")
		fmt.Println(argumetosSemNomePrograma[0])
		fmt.Println(argumetosSemNomePrograma[1])

	} else {
		fmt.Println("Informe pelo menos um argumento")
	}

}
