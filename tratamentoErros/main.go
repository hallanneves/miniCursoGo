package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

//Usuários é a strutura que contem um array de usuários
type Usuários struct {
	Usuários []Usuário `json:"users"`
}

// Usuário é a struct com os dados basse do usuário
// o nome do campo equivalente no json é especificado na declaraçãod da struct
type Usuário struct {
	Nome   string `json:"name"`
	Tipo   string `json:"type"`
	Idade  int    `json:"Age"`
	Social Social `json:"social"`
}

// Social é a struc que guarda as redes sociais
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	// Abre o arquivo
	jsonFile, err := os.Open("exemplo.json")
	// Se o arquivo não abrir retorna o erro
	if err != nil {
		//log Fatal sai do programa depois de mostrar o erro
		log.Fatal(err)
	}

	fmt.Println("Abriu o arquivo exemplo.json")
	// O defer, é umas das obras primas do GO.
	// ele é chamado antes do fim da função e antes de qualquer return
	// é possível ainda empilhar vários defer para serem executados antes do fim de uma função
	defer jsonFile.Close()

	// lê o arquivo como um array de bytes.
	// O que é o _ ???
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Inicializa o array de usuários
	var usuarios Usuários

	// o unmarshal converte o array de bytes na nossa estrutura
	// O que é esse & aí ??
	json.Unmarshal(byteValue, &usuarios)

	// itera os cadastros
	for i := 0; i < len(usuarios.Usuários); i++ {
		fmt.Println("Usuário Tipo: " + usuarios.Usuários[i].Tipo)
		fmt.Println("Usuário Idade: " + strconv.Itoa(usuarios.Usuários[i].Idade))
		fmt.Println("Usuário Nome: " + usuarios.Usuários[i].Nome)
		fmt.Println("Facebook Url: " + usuarios.Usuários[i].Social.Facebook)
	}

}
