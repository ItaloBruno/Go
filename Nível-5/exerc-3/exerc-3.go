/*
	Crie um novo tipo: veículo
	O tipo subjacente deve ser struct
	Deve conter os campos: portas, cor
	Crie dois novos tipos: caminhonete e sedan
	Os tipos subjacentes devem ser struct
	Ambos devem conter "veículo" como struct embutido
	O tipo caminhonete deve conter um campo bool chamado "traçãoNasQuatro"
	O tipo sedan deve conter um campo bool chamado "modeloLuxo"
	Usando os structs veículo, caminhonete e sedan:
	Usando composite literal, crie um valor de tipo caminhonete e dê valores a seus campos
	Usando composite literal, crie um valor de tipo sedan e dê valores a seus campos
	Demonstre estes valores.
	Demonstre um único campo de cada um dos dois.
*/

package main

import (
	"fmt"
)

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	traçãoNasQuatro bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {
	caminhonete1 := caminhonete{veiculo{2, "Azul"}, true}
	sedan1 := sedan{veiculo{4, "Cinza"}, false}

	fmt.Println("Dados da caminhote: ")
	fmt.Println(caminhonete1)

	fmt.Println("Dados do Sedan: ")
	fmt.Println(sedan1)

	fmt.Println("A caminhote tem tração nas quatro rodas? ", caminhonete1.traçãoNasQuatro)
	fmt.Println("O modelo do Sedan é de luxo? ", sedan1.modeloLuxo)
}
