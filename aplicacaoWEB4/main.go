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

// substituí as linhas repetidas por uma função onde tratamos todos os erros

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, err := template.ParseFiles(tmpl + ".html")
	//Erro ao carregar o template
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Erro em executar ou exibir o template
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//função que é chamada para responder a requisição do caminho /view/
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	//agora se entrarmos em uma página que não existe somos redirecionados para editar a página
	//Isso vai possibilitar a criarção de páginas novas para a nossa WIKI
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

//função que é chamada para responder a requisição do caminho /edit/
func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	//recebe a string do campo de name "body"
	body := r.FormValue("body")
	//Cria a estrutura da página
	p := &Page{Title: title, Body: []byte(body)}
	//Salva a página
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//Redireciona para a visualizaçao da págia
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	//Define os caminhos: qual URL é atendida por qual função
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
