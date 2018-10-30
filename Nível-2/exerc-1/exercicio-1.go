/*
	Escreva um programa que mostre um número em decimal, binário, e hexadecimal.
*/

package main

import (
	"fmt"
)

func main() {
	num := 30
	fmt.Println("Número: ", num)
	fmt.Printf("Decimal: %d | Binário: %b | Hexadecimal: %#x\n", num, num, num)
}
