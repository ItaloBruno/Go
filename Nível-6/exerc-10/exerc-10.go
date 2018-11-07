/*
	Demonstre o funcionamento de um closure.
	Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope.
*/

package main

import (
	"fmt"
)

func main() {
	x := incrementaContador()
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())

	y := incrementaContador()
	fmt.Println(y())
	fmt.Println(y())
	fmt.Println(y())
}

func incrementaContador() func() int {
	contador := 0
	return func() int {
		contador++
		return contador
	}
}
