package main

import (
	"fmt"
	"strconv"
)

type pessoaTipo struct {
	nome  string
	idade int
}

func imprimePessoa(pessoa pessoaTipo) {
	fmt.Println("Nome: " + pessoa.nome)
	fmt.Println("Idade: " + strconv.Itoa(pessoa.idade))
}

func main() {

	var arrayPessoa [3]pessoaTipo

	pessoa1 := pessoaTipo{"Gabriel", 17}
	arrayPessoa[0] = pessoa1

	var pessoa2 pessoaTipo
	pessoa2.nome = "Thiago"
	pessoa2.idade = 16

	arrayPessoa[1] = pessoa2

	arrayPessoa[2] = pessoaTipo{"Julia", 17}

	for _, p := range arrayPessoa {
		imprimePessoa(p)
	}

}
