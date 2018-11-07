/*
	Atribua uma função a uma variável.
	Chame essa função.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	f := func(x float64) float64 {
		aoQuadrado := math.Pow(x, 2.0)
		return aoQuadrado
	}
	fmt.Println("Resultado:", f(10))
}
