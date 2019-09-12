package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

//Page é a estrutura base que vamos usar para uma página
type Page struct {
	Title string

	Body []byte
}

// Função que são uma página em arquivo (nome do arquivo = Titulo.txt)
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

// Carrega uma página
func loadPage(title string) (*Page, error) {

	//Nome do arquivo que lê é recebido como parâmetro
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)

	//O go oferece suporte a vários parâmetros de saída
	if err != nil {
		return nil, err
	}

	return &Page{Title: title, Body: body}, nil

}

//função que é chamada para responder a requisição do caminho /view/
func viewHandler(w http.ResponseWriter, r *http.Request) {

	//Pega o nome do arquivo que vai carregar
	//como o r.URL.Path é um array de string, eu quero o que estiver depois do texto /view/
	// Exemplo /view/teste vai resultar tem title := "teste"
	title := r.URL.Path[len("/view/"):]

	p, err := loadPage(title)

	//Se o arquivo não existir vamos tratar o erro com essa mensagem por enquanto
	if err != nil {
		fmt.Fprintf(w, "<h1>Página não encontrada.</h1><div>%s</div>", err.Error())
		return
	}

	//Printa o HTML na variavel w
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func main() {

	//Define os caminhos: qual URL é atendida por qual função
	http.HandleFunc("/view/", viewHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))

}
