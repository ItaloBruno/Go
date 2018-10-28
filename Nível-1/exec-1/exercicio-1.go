/*
	Utilizando o operador curto de declaração, atribua estes valores às variáveis com os identificadores "x", "y", e "z".

    42
    "James Bond"
    true

    Agora demonstre os valores nestas variáveis, com:

    Uma única declaração print
    Múltiplas declarações print
*/

package main

import (
	"fmt"
)

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)

	fmt.Print("\nVariável x -> ")
	fmt.Printf("Valor: %v, Tipo: %T\n", x, x)

	fmt.Print("\nVariável y -> ")
	fmt.Printf("Valor: %v, Tipo: %T\n", y, y)

	fmt.Print("\nVariável z -> ")
	fmt.Printf("Valor: %v, Tipo: %T\n", z, z)
}
