/*
	Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
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

	dadosPessoais["Paula_Noronha"] = []string{
		"Escrever",
		"Ler Hqs",
	}

	for nome, hobbies := range dadosPessoais {
		fmt.Printf("Nome da Pessoa: %v\t", nome)
		for _, valor := range hobbies {
			fmt.Printf("Hobbie: %v\t", valor)
		}
		fmt.Println()
	}

	fmt.Println("Excluindo um elemento do map....")
	delete(dadosPessoais, "João_Bruno")

	for nome, hobbies := range dadosPessoais {
		fmt.Printf("Nome da Pessoa: %v\t", nome)
		for _, valor := range hobbies {
			fmt.Printf("Hobbie: %v\t", valor)
		}
		fmt.Println()
	}
}
