/*
	Crie uma dadosPessoais contendo dadosPessoaiss de strings ([][]string).
	Atribua valores a este dadosPessoais multi-dimensional da seguinte maneira:
	"Nome", "Sobrenome", "Hobby favorito"
	Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
*/

package main

import (
	"fmt"
)

func main() {
	dadosPessoais := [][]string{
		[]string{
			"João",
			"Bruno",
			"Programar",
		},
		[]string{
			"Maria",
			"Fátima",
			"Reclamar da vida",
		},
		[]string{
			"Laís",
			"Farias",
			"Ler Livros",
		},
	}
	for _, valor := range dadosPessoais {
		fmt.Printf("Nome: %v - Sobrenone: %v - Hobby favorito: %v\n", valor[0], valor[1], valor[2])
	}
}
