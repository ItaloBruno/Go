/*
	Crie e utilize uma função anônima.
*/

package main

import (
	"fmt"
)

func main() {
	slice := []int{100, 200}

	func(sliceDeInteiros ...int) {
		resultadoSoma := 0
		for _, valor := range sliceDeInteiros {
			resultadoSoma += valor
		}
		fmt.Println("A soma dos elementos da slice é:", resultadoSoma)
	}(slice...)
}
