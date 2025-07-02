package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero) // cria uma cópia da variável

	fmt.Println(numeroInvertido)

	novoNumero := 40
	inverterSinalPonteiro(&novoNumero) // envia o endereço de memória da variável interferindo direto no valor dela
	fmt.Println(novoNumero)

}
