package main

import "fmt"

func main() {
	canal := make(chan string, 2) // só bloqueia o trafego de dados quando o canal atigir a capacidade máxima

	canal <- "Olá mundo"
	canal <- "Programa em go"

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
