/*
	Usando uma literal composta:
	Crie um array que suporte 5 valores to tipo int
	Atribua valores aos seus índices
	Utilize range e demonstre os valores do array.
	Utilizando format printing, demonstre o tipo do array.
*/

package main

import (
	"fmt"
)

func main() {
	listaDeInteiros := [5]int{
		1, 2, 3, 4, 5}
	fmt.Println(listaDeInteiros)
	for indice, valor := range listaDeInteiros {
		fmt.Printf("Índice: %v - Valor: %v\n", indice, valor)
	}
	fmt.Printf("Tipo do meu array: %T\n", listaDeInteiros)
}
