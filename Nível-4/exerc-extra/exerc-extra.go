/*
	Exercício extra: tente acessar todos os itens de uma slice sem utilizar range
*/

package main

import (
	"fmt"
)

func main() {
	sabores := []string{
		"morango",
		"chocolate",
		"baunilha",
		"ovo maltine",
		"oreo"}
	fmt.Println(sabores)

	sorvete := sabores[:]
	fmt.Println(sorvete)
}
