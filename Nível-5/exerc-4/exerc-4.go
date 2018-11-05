/*
	Crie e use um struct anônimo.
	Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
*/

package main

import (
	"fmt"
)

func main() {
	pessoa := struct {
		nomeSobrenome map[string]string
		idade         int
		cores         []string
	}{nomeSobrenome: map[string]string{
		"Drew": "Brees",
	},
		idade: 39,
		cores: []string{
			"Branco",
			"Dourado",
			"Preto",
		}}
	fmt.Println("Nome e Sobrenome | Idade | Cores: ")
	fmt.Println(pessoa)
}
