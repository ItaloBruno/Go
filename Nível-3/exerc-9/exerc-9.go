/*
	Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
*/

package main

import (
	"fmt"
)

func main() {
	esporteFavorito := "basquete"
	switch esporteFavorito {
	case "futebol":
		fmt.Println("FUTEBOOOOOOL")
	case "basquete":
		fmt.Println("BASQUETE!!!!!!")
	default:
		fmt.Println("Não gosto de esporte!")
	}
}
