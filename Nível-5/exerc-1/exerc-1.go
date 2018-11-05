/*
	Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
	Nome
	Sobrenome
	Sabores favoritos de sorvete
	Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
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
	pessoa1 := pessoa{
		nome:      "tulio",
		sobrenome: "maravilha",
		sabores: []string{
			"Baunilha",
			"Chocolate",
		},
	}

	pessoa2 := pessoa{"Renato", "Gaúcho", []string{
		"Morango",
		"Creme",
	}}

	imprimiPessoa(pessoa1)
	fmt.Println()
	imprimiPessoa(pessoa2)

}

func imprimiPessoa(pessoa pessoa) {
	fmt.Println("Nome:", pessoa.nome)
	fmt.Println("Sobrenome:", pessoa.sobrenome)
	fmt.Println("Sabores de sorvete que eu gosto:")

	for _, sabor := range pessoa.sabores {
		fmt.Println(sabor)
	}
}
