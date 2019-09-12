package main

import (
	"fmt"

	pacoteContatos "github.com/hallanneves/miniCursoGo/aplicacaoWEB/contatos"
)

var contatos []pacoteContatos.Contato

func main() {

	contatos = pacoteContatos.CarregaContatos().Contatos

	fmt.Println(contatos)

	contatoNovo := pacoteContatos.Contato{Nome: "Hallan", Idade: 28, Telefone: "(53) 3232-3232"}

	contatos = append(contatos, contatoNovo)

	contatosPersistir := pacoteContatos.Contatos{Contatos: contatos}

	pacoteContatos.SalvaContatos(contatosPersistir)

}
