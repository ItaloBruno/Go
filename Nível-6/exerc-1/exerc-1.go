/*
	Crie uma função que retorne um int
	Crie outra função que retorne um int e uma string
	Chame as duas funções
	Demonstre seus resultados
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("2*4:", dobrarNúmero(4))
	soma, msg := somaMaisDez(5)
	fmt.Printf("%v: %v\n", msg, soma)
}

func dobrarNúmero(num int) int {
	return num * 2
}

func somaMaisDez(num int) (int, string) {
	return num + 10, "Resultado"
}
