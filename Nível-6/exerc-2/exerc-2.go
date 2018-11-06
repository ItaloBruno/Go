/*
	Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
	Passe um valor do tipo slice of int como argumento para a função;
	Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
	Passe um valor do tipo slice of int como argumento para a função.
*/

package main

import (
	"fmt"
)

func main() {
	slice := []int{100, 200}
	resultado := soma(slice...)
	fmt.Println("Resultado da soma dos elementos do slice:", resultado)
	resultado = somaSlice(slice)
	fmt.Println("Resultado da soma dos elementos do slice:", resultado)
}

func soma(sliceDeInteiros ...int) int {
	resultadoSoma := 0
	for _, valor := range sliceDeInteiros {
		resultadoSoma += valor
	}
	return resultadoSoma
}

func somaSlice(sliceDeInteiros []int) int {
	resultadoSoma := 0
	for _, valor := range sliceDeInteiros {
		resultadoSoma += valor
	}
	return resultadoSoma
}
