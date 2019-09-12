package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"text/template"

	pacoteContatos "github.com/hallanneves/miniCursoGo/aplicacaoWEB7/contatos"
)

var contatos []pacoteContatos.Contato

var templates = template.Must(template.ParseFiles("novo.html", "lista.html"))

var validPath = regexp.MustCompile("^/(lista|save|novo)")

//* Adapta a função para somente executar um temmplate e recebe o nome dele como parâmetro
// substituí as linhas repetidas por uma função onde tratamos todos os erros
func renderTemplate(w http.ResponseWriter, tmpl string, atributos ...interface{}) {
	err := templates.ExecuteTemplate(w, tmpl+".html", atributos)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//!MELHORIA FINAL ULTRA MASTER GO
//* No go podemos fazer uma função que recebe outra função como parâmetro e assim podemos
//* executar um procedimento comum para todas as funções e depois um especifico, a função
//* que recebemos como parâmetro

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r)
	}
}

//* Alterada para usar o renderTemplate e o getTitle
//função que é chamada para responder a requisição do caminho /view/
func viewHandler(w http.ResponseWriter, r *http.Request) {
	var contatos, err = pacoteContatos.CarregaContatos()
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	renderTemplate(w, "lista", contatos)
}

//* Alterada para usar o renderTemplate e o getTitle
//função que é chamada para responder a requisição do caminho /edit/
func editHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "novo")
}

//* Alterada para usar o renderTemplate e o getTitle
func saveHandler(w http.ResponseWriter, r *http.Request) {
	nome := r.FormValue("nome")
	idade, err := strconv.Atoi(r.FormValue("idade"))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	telefone := r.FormValue("telefone")

	contatoNovo := pacoteContatos.Contato{Nome: nome, Idade: idade, Telefone: telefone}
	contatos = append(contatos, contatoNovo)
	contatoEstrutura := pacoteContatos.Contatos{Contatos: contatos}

	pacoteContatos.SalvaContatos(contatoEstrutura)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/lista/", http.StatusFound)
}

func main() {

	contatosEstrutura, err := pacoteContatos.CarregaContatos()
	if err != nil {
		fmt.Println(err)
		return
	}
	contatos = contatosEstrutura.Contatos

	http.HandleFunc("/lista/", makeHandler(viewHandler))
	http.HandleFunc("/novo/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
