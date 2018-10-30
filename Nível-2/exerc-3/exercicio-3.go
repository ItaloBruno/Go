/*
	Crie constantes tipadas e não-tipadas.
	Demonstre seus valores.
*/

package main

import (
	"fmt"
)

const (
	a         = 10
	b float64 = 20.0
)

func main() {
	fmt.Printf("const a: %v %T\n", a, a)
	fmt.Printf("const b: %v %T\n", b, b)
}
