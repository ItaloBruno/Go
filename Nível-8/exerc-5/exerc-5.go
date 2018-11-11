/*
	Partindo do código abaixo, ordene os []user por idade e sobrenome.
	https://play.golang.org/p/BVRZTdlUZ_
	Os valores no campo Sayings devem ser ordenados tambem, e demonstrados de maneira esteticamente harmoniosa.
*/

package main

import (
	"fmt"
	"sort"
)

type user struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

type ordenaPorIdade []user

func (u ordenaPorIdade) Len() int           { return len(u) }
func (u ordenaPorIdade) Less(i, j int) bool { return u[i].Age < u[j].Age }
func (u ordenaPorIdade) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

type ordenaPorSobrenome []user

func (u ordenaPorSobrenome) Len() int           { return len(u) }
func (u ordenaPorSobrenome) Less(i, j int) bool { return u[i].Last < u[j].Last }
func (u ordenaPorSobrenome) Swap(i, j int)      { u[i], u[j] = u[j], u[i] }

func main() {
	u1 := user{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}

	u2 := user{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := user{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []user{u1, u2, u3}
	for _, v := range users {
		sort.Strings(v.Sayings)
	}

	fmt.Println("Usuários não ordenados:")
	imprimirUser(users)

	sort.Sort(ordenaPorIdade(users))
	fmt.Println("\n\nOrdenado por Idade:")
	imprimirUser(users)

	fmt.Println()

	sort.Sort(ordenaPorSobrenome(users))

	fmt.Println("\n\nOrdenado por Sobrenome:")
	imprimirUser(users)
}

func imprimirUser(sliceUsers []user) {
	for _, v := range sliceUsers {
		fmt.Printf("Name: %v - ", v.First)
		fmt.Printf("Last: %v - ", v.Last)
		fmt.Printf("Age: %v\n", v.Age)
		fmt.Println("Sayings:")
		for j, value := range v.Sayings {
			fmt.Printf("%v - %v\n", j+1, value)
		}
		fmt.Println()
	}
}
