/*
	Callback: passe uma função como argumento a outra função.
*/

package main

import (
	"fmt"
)

func main() {
	totalDaSoma := somenteNúmerosImpares(somarNúmeros, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println("Total da soma dos números do slice e que são ímpares:", totalDaSoma)
}

func somarNúmeros(x ...int) int {
	soma := 0
	for _, valor := range x {
		soma += valor
	}
	return soma
}

func somenteNúmerosImpares(somarNúmeros func(x ...int) int, y ...int) int {
	var númerosImpares []int
	for _, valor := range y {
		if valor%2 != 0 {
			númerosImpares = append(númerosImpares, valor)
		}
	}
	valorTotal := somarNúmeros(númerosImpares...)
	return valorTotal
}
