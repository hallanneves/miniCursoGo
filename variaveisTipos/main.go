package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Declaração de um variável

	var i int = 1

	var j = 2

	k := 3

	//Uso e impressão na tela
	soma := i + j + k

	fmt.Println("Soma = " + strconv.Itoa(soma))

	//Arrays
	var testeArray [1]string
	testeArray[0] = "testeArray"

	fmt.Println(testeArray)

	//Slice
	var palavras []string

	palavras = append(palavras, "Teste")
	palavras = append(palavras, "Olá")
	palavras = append(palavras, "Mundo")

	palavrasSemLixo := palavras[1:]

	fmt.Println(palavrasSemLixo[0] + " " + palavrasSemLixo[1])

	//Array com elementos na linha

	palavrasDiferentes := [2]string{"Mundo", "Olá"}

	fmt.Println(palavrasDiferentes)

	// Maps

	var m map[string]int
	m = make(map[string]int)

	m["route"] = 66

	fmt.Println("Rota: " + strconv.Itoa(m["route"]))

	//Não façam isso, ou façam, vai de vocês, eu acho estranho

	var coração = 4
	fmt.Println(coração)
}
