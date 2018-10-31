/*
	Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
*/

package main

import (
	"fmt"
)

func main() {
	for num := 10; num <= 100; num++ {
		fmt.Printf("Resto da divisão de %v por 4 = %v\n", num, num%4)
	}
}
