package main

import (
	"errors"
	"fmt"
)

func main() {
	// NUMEROS INTEIROS
	var numero int64 = 1000000000
	fmt.Println(numero)

	var numero2 uint32 = 10000
	fmt.Println(numero2)

	//Alias
	// INT32 = RUNE
	var numero3 rune = 123456
	fmt.Println(numero3)

	// BYTE = UINT8
	var numero4 byte = 123
	fmt.Println(numero4)

	// FIM NUMEROS INTEIROS

	// NUMEROS REAIS

	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)

	var numeroReal2 float64 = 123000000.45
	fmt.Println(numeroReal2)

	// FIM NUMEROS REAIS

	// STRINGS
	var str string = "Texto"
	fmt.Println(str)

	char := '8' // retorna apenas a posição inteira do char na tabela ascii
	fmt.Println(char)

	// FIM STRINGS

	var booleano bool
	fmt.Println(booleano)

	var erro error = errors.New("erro interno")
	fmt.Println(erro)

}
