/*
	Crie um tipo. O tipo subjacente deve ser int.
    Crie uma variável para este tipo, com o identificador "x", utilizando a palavra-chave var.
    Na função main:

    Demonstre o valor da variável "x"
    Demonstre o tipo da variável "x"
    Atribua 42 à variável "x" utilizando o operador "="
    Demonstre o valor da variável "x"

    Para os aventureiros: https://golang.org/ref/spec#Types
*/

package main

import (
	"fmt"
)

type rams int

var x rams

func main() {
	fmt.Printf("\nValor de x: %v\n", x)
	fmt.Printf("\nTipo de x: %T\n", x)

	x = 42

	fmt.Printf("\nValor de x: %v\n", x)
}
