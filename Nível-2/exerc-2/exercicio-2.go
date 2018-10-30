/*
	Escreva expressões utilizando os seguintes operadores, e atribua seus valores a variáveis.
	==
	!=
	<=
	<
	>=
	>
	Demonstre estes valores.
*/

package main

import (
	"fmt"
)

func main() {
	x := 10
	y := 10
	result1 := x < y
	result2 := x > y
	result3 := x <= y
	result4 := x >= y
	result5 := x == y
	result6 := x != y

	fmt.Printf("%v < %v? %v\n", x, y, result1)
	fmt.Printf("%v > %v? %v\n", x, y, result2)
	fmt.Printf("%v <= %v? %v\n", x, y, result3)
	fmt.Printf("%v >= %v? %v\n", x, y, result4)
	fmt.Printf("%v == %v? %v\n", x, y, result5)
	fmt.Printf("%v != %v? %v\n", x, y, result6)

}
