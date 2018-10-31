/*
	Crie um programa que utilize a declaração switch, sem switch statement especificado.
*/

package main

import (
	"fmt"
)

func main() {
	turno := "manha"
	switch {
	case turno == "manha":
		fmt.Println("É de manhã")
	case turno == "tarde":
		fmt.Println("É de tarde")
	default:
		fmt.Println("É hora de dormir")
	}
}
