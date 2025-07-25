package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var var1 int = 10
	var var2 int = var1

	fmt.Println(var1, var2)

	var1++

	fmt.Println(var1, var2)

	// PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	var var3 int = 100
	var ponteiro *int

	ponteiro = &var3

	fmt.Println(var3, ponteiro)
	fmt.Println(var3, *ponteiro) //desreferenciação do ponteiro

	var3 = 150

	fmt.Println(var3, ponteiro)
	fmt.Println(var3, *ponteiro) //desreferenciação do ponteiro
}
