/*
	Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
	Por exemplo: 65 U+0041 'A' U+0041 'A' U+0041 'A' 66 U+0042 'B' U+0042 'B' U+0042 'B' ...e por aí vai.
*/

package main

import (
	"fmt"
)

func main() {
	for caracter := 65; caracter < 91; caracter++ {
		fmt.Printf("\n%v \n", caracter)
		fmt.Printf("\t%#U \n\t%#U \n\t%#U \n", caracter, caracter, caracter)
	}
}
