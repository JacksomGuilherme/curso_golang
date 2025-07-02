package main

import (
	"fmt"
)

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var user1 usuario
	user1.nome = "Davi"
	user1.idade = 16
	fmt.Println(user1)

	endereco1 := endereco{"Rua dos bobos", 0}

	user2 := usuario{"João", 24, endereco1}
	fmt.Println(user2)

	endereco2 := endereco{"Avenida Moaci", 57}

	var user3 usuario = usuario{"Jacksom", 26, endereco2}
	fmt.Println(user3)

	user4 := usuario{nome: "Brian"}
	fmt.Println(user4)

	user5 := usuario{idade: 30}
	fmt.Println(user5)
}
