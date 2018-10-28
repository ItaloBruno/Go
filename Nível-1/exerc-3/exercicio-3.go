/*
	Utilizando a solução do exercício anterior:

    Em package-level scope, atribua os seguintes valores às variáveis:
    para "x" atribua 42
    para "y" atribua "James Bond"
    para "z" atribua true
    Na função main:
    Use fmt.Sprintf para atribuir todos esses valores a uma única variável. Faça essa atribuição de tipo string a uma variável de nome "s" utilizando o operador curto de declaração.
    Demonstre a variável "s".
*/

package main

import (
	"fmt"
)

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("x: %v, y: %v, z: %v", x, y, z)
	fmt.Println(s)
}
