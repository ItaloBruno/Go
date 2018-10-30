/*
	Crie um programa que:
	Atribua um valor int a uma variável
	Demonstre este valor em decimal, binário e hexadecimal
	Desloque os bits dessa variável 1 para a esquerda, e atribua o resultado a outra variável
	Demonstre esta outra variável em decimal, binário e hexadecimal
*/
package main

import (
	"fmt"
)

func main() {
	numero := 200
	fmt.Println("Número: ", numero)
	fmt.Printf("Decimal: %d | Binário: %b | Hexadecimal: %#x\n", numero, numero, numero)

	fmt.Printf("\nDeslocando 1 bit para a esquerda...\n\n")
	numero_modificado := numero << 1
	fmt.Println("Número: ", numero_modificado)
	fmt.Printf("Decimal: %d | Binário: %b | Hexadecimal: %#x\n", numero_modificado, numero_modificado, numero_modificado)
}
