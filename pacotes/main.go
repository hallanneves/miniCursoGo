package main

import (
	pessoaPacote "github.com/hallanneves/miniCursoGo/pacotes/pessoa"
)

func main() {

	var arrayPessoa [3]pessoaPacote.Pessoa

	// a inserção pode e deve ser implicita para facilitar a legibilidade do código no futuro
	pessoa1 := pessoaPacote.Pessoa{Nome: "Gabriel", Idade: 17}
	arrayPessoa[0] = pessoa1

	arrayPessoa[1] = pessoaPacote.Pessoa{Nome: "Thiago", Idade: 16}
	arrayPessoa[2] = pessoaPacote.Pessoa{Nome: "Julia", Idade: 17}

	for _, p := range arrayPessoa {
		p.ImprimeExterno()
	}

	for _, p := range pessoaPacote.PessoasPublicas {
		pessoaPacote.ImprimeComumDopacote(p)
	}

}
