package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("")

	p1 := pessoa{"João", "Pedro", 20, 178}
	fmt.Println(p1)

	e1 := estudante{p1, "ADS", "FATEC"}
	fmt.Println(e1)
	fmt.Println("----------------")
	fmt.Println(e1.nome)
	fmt.Println(e1.idade)
	fmt.Println(e1.curso)

}
