package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Variaticas")

	totalDaSoma := soma(1, 2, 3, 0)

	fmt.Println(totalDaSoma)

	escrever("Ganhadores: ", 55, 23, 65, 85, 60)
}
