/*
	Anote (à mão) o resultado das expressões:
	fmt.Println(true && true)
	fmt.Println(true && false)
	fmt.Println(true || true)
	fmt.Println(true || false)
	fmt.Println(!true)
*/

package main

import (
	"fmt"
)

func main() {
	fmt.Println("true && true...")
	fmt.Printf("Resposta: %v\n\n", true && true)
	fmt.Println("true && false...")
	fmt.Printf("Resposta: %v\n\n", true && false)
	fmt.Println("true || true...")
	fmt.Printf("Resposta: %v\n\n", true || true)
	fmt.Println("true || false...")
	fmt.Printf("Resposta: %v\n\n", true && false)
	fmt.Println("!true...")
	fmt.Printf("Resposta: %v\n\n", !true)
}
