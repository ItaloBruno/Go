/*
	Crie um tipo "quadrado"
	Crie um tipo "círculo"
	Crie um método "área" para cada tipo que calcule e retorne a área da figura
	Área do círculo: 2 * π * raio
	Área do quadrado: lado * lado
	Crie um tipo "figura" que defina como interface qualquer tipo que tiver o método "área"
	Crie uma função "info" que receba um tipo "figura" e retorne a área da figura
	Crie um valor de tipo "quadrado"
	Crie um valor de tipo "círculo"
	Use a função "info" para demonstrar a área do "quadrado"
	Use a função "info" para demonstrar a área do "círculo"
*/

package main

import (
	"fmt"
	"math"
)

type quadrado struct {
	lado float64
}

type círculo struct {
	raio float64
}

func (q quadrado) área() {
	resultado := math.Pow(q.lado, 2.0)
	fmt.Println("Área do quadrado:", resultado)
}

func (c círculo) área() {
	resultado := c.raio * 2 * math.Pi
	fmt.Println("Área do círculo:", resultado)
}

type figura interface {
	área()
}

func info(f figura) {
	f.área()
}

func main() {
	quadrado1 := quadrado{15.0}
	círculo1 := círculo{5.25}

	info(quadrado1)
	info(círculo1)
}
