/*
	Crie um programa que demonstre o funcionamento da declaração if.
*/

package main

import (
	"fmt"
)

func main() {
	for numero := 0; numero <= 100; numero++ {
		if numero%2 == 0 {
			fmt.Printf("%v é par\n", numero)
		} else {
			fmt.Printf("%v é impar\n", numero)
		}
	}
}
