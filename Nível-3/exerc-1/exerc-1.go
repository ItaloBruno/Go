// Põe na tela: todos os números de 1 a 10000.

package main

import (
	"fmt"
)

func main() {
	for num := 1; num <= 10000; num++ {
		fmt.Printf("Número: %v\n", num)
	}
}
