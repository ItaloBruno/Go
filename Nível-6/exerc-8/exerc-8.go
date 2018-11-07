/*
	Crie uma função que retorna uma função.
	Atribua a função retornada a uma variável.
	Chame a função retornada.
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	f := retornaNúmeroAoCubo()
	fmt.Println("2 elevado ao cubo é:", f(2.0))
}

func retornaNúmeroAoCubo() func(num float64) float64 {
	return func(num float64) float64 {
		return math.Pow(num, 3.0)
	}
}
