package contatos

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//Contatos é a strutura que contem um array de Contatos
type Contatos struct {
	Contatos []Contato `json:"contatos"`
}

// Contato é a struct com os dados basse do Contato
// o nome do campo equivalente no json é especificado na declaraçãod da struct
type Contato struct {
	Nome     string `json:"nome"`
	Idade    int    `json:"idade"`
	Telefone string `json:"telefone"`
}

//CarregaContatos do json
func CarregaContatos() (Contatos, error) {
	var contatos Contatos

	jsonFile, err := os.Open("contatos.json")
	if err != nil {
		return contatos, err
	}
	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return contatos, err
	}
	json.Unmarshal(byteValue, &contatos)

	return contatos, nil
}

//SalvaContatos persiste os contatos adicionados
func SalvaContatos(contatos Contatos) error {
	file, err := json.MarshalIndent(contatos, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile("test.json", file, 0644)

	return err
}
