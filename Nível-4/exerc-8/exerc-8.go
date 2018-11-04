/*
	Crie um map com key tipo string e value tipo []string.
	Key deve conter nomes no formato sobrenome_nome
	Value deve conter os hobbies favoritos da pessoa
	Demonstre todos esses valores e seus índices.
*/

package main

import (
	"fmt"
)

func main() {
	dadosPessoais := map[string][]string{
		"João_Bruno": []string{
			"Programar",
			"Jogar Video Game",
		},
		"Maria_Fátima": []string{
			"Reclamar da vida",
			"Andar por aí",
		},
		"Laís_Farias": []string{
			"Ler Livros",
			"Contruir móveis",
		},
	}
	fmt.Println(dadosPessoais)

	for nome, hobbies := range dadosPessoais {
		fmt.Printf("Nome da Pessoa: %v\t", nome)
		for _, valor := range hobbies {
			fmt.Printf("Hobbie: %v\t", valor)
		}
		fmt.Println()
	}
}
