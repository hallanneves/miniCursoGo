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

package main

import (
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"text/template"
)

var templates = template.Must(template.ParseFiles("novo.html", "lsita.html"))

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

//!MELHORIA FINAL ULTRA MASTER GO
//* No go podemos fazer uma função que recebe outra função como parâmetro e assim podemos
//* executar um procedimento comum para todas as funções e depois um especifico, a função
//* que recebemos como parâmetro

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

//* Alterada para usar o renderTemplate e o getTitle
//função que é chamada para responder a requisição do caminho /view/
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

//* Alterada para usar o renderTemplate e o getTitle
//função que é chamada para responder a requisição do caminho /edit/
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

//* Alterada para usar o renderTemplate e o getTitle
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func main() {
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))

	log.Fatal(http.ListenAndServe(":8080", nil))
}
