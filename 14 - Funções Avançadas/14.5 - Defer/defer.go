package main

import "fmt"

func f1() {
	fmt.Println("Executando func 1")
}

func f2() {
	fmt.Println("Executando func 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média Calculada. Resultado será printado")
	fmt.Println("Validando aluno aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	// defer f1()
	// f2()

	fmt.Println(alunoAprovado(6, 5))
}
