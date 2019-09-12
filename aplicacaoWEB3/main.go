package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"text/template"
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
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	//O go por padrão trabalha co templates WEB. Podemos então fazer páginas
	// onde vamos colocar o nosso HTML
	//utilizando o template não vamos ter erro caso o p retorne nil
	// a página vai estar vazia
	t, _ := template.ParseFiles("view.html")
	t.Execute(w, p)
}

//função que é chamada para responder a requisição do caminho /edit/
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	//Não estamos tratando o erro de templete inesistente, depois vamos tratar
	t, _ := template.ParseFiles("edit.html")
	t.Execute(w, p)
}

func main() {
	//Define os caminhos: qual URL é atendida por qual função
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	//http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
