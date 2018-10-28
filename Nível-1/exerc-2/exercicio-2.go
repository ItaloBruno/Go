/*
	Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis:

    Identificador "x" deverá ter tipo int
    Identificador "y" deverá ter tipo string
    Identificador "z" deverá ter tipo bool

    Na função main:

    Demonstre os valores de cada identificador
    O compilador atribuiu valores para essas variáveis. Como esses valores se chamam?
*/

package main

import (
	"fmt"
)

var x int
var y string
var z bool

func main() {

	fmt.Print("\nVariável x -> ")
	fmt.Printf("Valor zero: %v, Tipo: %T\n", x, x)

	fmt.Print("\nVariável y -> ")
	fmt.Printf("Valor zero: %v, Tipo: %T\n", y, y)

	fmt.Print("\nVariável z -> ")
	fmt.Printf("Valor zero: %v, Tipo: %T\n", z, z)

	fmt.Println("\nValor zero é o valor que o compilador atribue à variável que não foi inicializada")
}
