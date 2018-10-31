/*
	Utilizando a solução anterior, adicione as opções else if e else.
*/

package main

import (
	"fmt"
)

func main() {
	for numero := 0; numero <= 100; numero++ {
		if numero%2 == 0 {
			fmt.Printf("%v é par\n", numero)
		} else if numero%2 != 0 {
			fmt.Printf("%v é impar\n", numero)
		} else {
			fmt.Printf("%v é nenhum dos dois\n", numero)
		}
	}
}
