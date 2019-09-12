package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
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

//! MELHORIA 1
//* Declarar os templates faz com que carreguemos eles somente uma vez em memória
//* O must é um "empacotador" que serve para caso um templete não consiga ser carregado dar um fatal error
//* O Parse files carrega os templates, usa como referencia para chamadas o próprio nome do template
var templates = template.Must(template.ParseFiles("edit.html", "view.html"))

//!Melhoria 2
//* O sitema está aceitando todo tipo de url, sendo que só temos 3 tratadas
//* O go suporta fazer expressões regulares para validar textos
//* Vamos usar isso para melhorar nosso código
var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

//* Criando uma função para retornar o titulo da página, que é utilizada nos handlers
//* Podemos validar a nossa URL de entrada
func getTitle(w http.ResponseWriter, r *http.Request) (string, error) {
	//Verifica se a string da math com a expressão
	m := validPath.FindStringSubmatch(r.URL.Path)
	//Se não da é porque não temos está página
	if m == nil {
		http.NotFound(w, r)
		return "", errors.New("Invalid Page Title")
	}
	return m[2], nil // The title is the second subexpression.
}

//* Adapta a função para somente executar um temmplate e recebe o nome dele como parâmetro
// substituí as linhas repetidas por uma função onde tratamos todos os erros
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//* Alterada para usar o renderTemplate e o getTitle
//função que é chamada para responder a requisição do caminho /view/
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

//* Alterada para usar o renderTemplate e o getTitle
//função que é chamada para responder a requisição do caminho /edit/
func editHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

//* Alterada para usar o renderTemplate e o getTitle
func saveHandler(w http.ResponseWriter, r *http.Request) {
	title, err := getTitle(w, r)
	if err != nil {
		return
	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err = p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	//Define os caminhos: qual URL é atendida por qual função
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
