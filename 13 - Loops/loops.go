package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j++ {
	// 	fmt.Println("Incrementando i", j)
	// 	time.Sleep(time.Second)
	// }

	nomes := [3]string{"Jack", "Iza", "Lino"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	fmt.Println("---------------------------")

	for _, nome := range nomes {
		fmt.Println(nome)
	}

	fmt.Println("---------------------------")

	for indice, letra := range "JACKSOM" {
		fmt.Println(indice, string(letra))
	}

	fmt.Println("---------------------------")

	usuario := map[string]string{
		"nome":      "Jacksom",
		"sobrenome": "Guilherme",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	type usuarioStruct struct {
		nome      string
		sobrenome string
		idade     int
	}

	usuarios2 := [2]usuarioStruct{{"Jacksom", "Guilherme", 26}, {"Izabelly", "Cristine", 26}}

	for indice, usuario := range usuarios2 {
		fmt.Println(indice, usuario)
	}

}
