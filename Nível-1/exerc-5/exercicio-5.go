/*
    Utilizando a solução do exercício anterior:

    Em package-level scope, utilizando a palavra-chave var, crie uma variável com o identificador "y". O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.
    Na função main:
    Isto já deve estar feito:
    Demonstre o valor da variável "x"
    Demonstre o tipo da variável "x"
    Atribua 42 à variável "x" utilizando o operador "="
    Demonstre o valor da variável "x"
    Agora faça tambem:
    Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando o operador "=", atribua o valor de "x" a "y"
    Demonstre o valor de "y"
	Demonstre o tipo de "y"
*/

package main

import (
	"fmt"
)

type rams int

var x rams
var y int

func main() {
	fmt.Printf("\nValor de x: %v\n", x)
	fmt.Printf("\nTipo de x: %T\n", x)

	x = 42

	fmt.Printf("\nValor de x: %v\n", x)

	fmt.Println("Convertendo X em Y...")
	y = int(x)
	fmt.Printf("\nValor de y: %v\n", y)
	fmt.Printf("\nTipo de y: %T\n", y)
}
