/*
	Crie um loop utilizando a sintaxe: for {}
	Utilize-o para demonstrar os anos desde que você nasceu.
*/

package main

import (
	"fmt"
)

const (
	anoAtual = 2018
)

func main() {
	anos := 1994
	for {
		fmt.Printf("Ano: %v\n", anos)
		anos++
		if anos > anoAtual {
			break
		}
	}
}
