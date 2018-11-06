/*
	Utilize a declaração defer de maneira que demonstre que sua execução só ocorre ao final do contexto ao qual ela pertence.
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Exemplificação do uso do defer...")
	defer fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")
	defer fmt.Println("5")
	fmt.Println("6")
	fmt.Println("7")
	fmt.Println("8")
	fmt.Println("9")
	fmt.Println("10")
}
