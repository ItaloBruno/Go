/*
	Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
	Demonstre estes valores.
*/

package main

import (
	"fmt"
)

const (
	ano1 = 2019 + iota
	ano2
	ano3
	ano4
)

func main() {

	fmt.Println("Os próximos 4 anos...")
	fmt.Println(ano1)
	fmt.Println(ano2)
	fmt.Println(ano3)
	fmt.Println(ano4)
}
