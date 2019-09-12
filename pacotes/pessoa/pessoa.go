package pessoa

import "fmt"

//Pessoa precisa ser comentado porque é uma extrutura visivel de fora do pacote
type Pessoa struct {
	Nome  string
	Idade int
	//campo que só pode ser acessado de dentro do pacote
	campoLocal int
}

//variável local do pacote
var pessoas = [3]Pessoa{Pessoa{"André", 11, 1}, Pessoa{"Julio", 12, 2}, Pessoa{"Romário", 11, 3}}

//PessoasPublicas é uma variável global e precisa do comentário iniciando com o seu nome
var PessoasPublicas = pessoas[1:]

func imprimeInterno(p Pessoa) {
	fmt.Println(p)
}

//ImprimeExterno também é aberto par outros pacotes e é uma funcao de pessoa
func (p Pessoa) ImprimeExterno() {
	imprimeInterno(p)
}

//ImprimeComumDopacote função aberta do pacote
func ImprimeComumDopacote(p Pessoa) {
	fmt.Println(p)
}
