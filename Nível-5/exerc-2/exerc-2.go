/*
	Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
	Demonstre os valores do map utilizando range.
	Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
*/

package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	mapaDePessoas := make(map[string]pessoa)
	mapaDePessoas["maravilha"] = pessoa{
		nome:      "tulio",
		sobrenome: "maravilha",
		sabores: []string{
			"Baunilha",
			"Chocolate",
		},
	}

	mapaDePessoas["Gaúcho"] = pessoa{"Renato", "Gaúcho", []string{
		"Morango",
		"Creme",
	}}

	for _, valor := range mapaDePessoas {
		fmt.Println("Nome:", valor.nome)
		fmt.Println("Sobrenome:", valor.sobrenome)
		fmt.Println("Sabores de sorvete que eu gosto:")

		for _, sabor := range valor.sabores {
			fmt.Println(sabor)
		}
		fmt.Println()
	}
}
