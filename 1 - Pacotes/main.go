package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Executando do main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("132")
	fmt.Println(erro)
}
