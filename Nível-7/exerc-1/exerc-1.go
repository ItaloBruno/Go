/*
	Crie um valor e atribua-o a uma variável.
	Demonstre o endereço deste valor na memória.
*/

package main

import (
	"fmt"
)

func main() {
	x := 2000
	y := &x
	fmt.Println("O endereço da variável x é:", y)
	fmt.Println("O valor armazenado neste endereço é:", *y)
}
