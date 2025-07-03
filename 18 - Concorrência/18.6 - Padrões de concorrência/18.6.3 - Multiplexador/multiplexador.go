package main

import (
	"fmt"
	"math/rand"
	"time"
)

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}

func multiplexar(canalDeEntrada1, canaldeEntrada2 <-chan string) <-chan string {
	canaldeSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalDeEntrada1:
				canaldeSaida <- mensagem
			case mensagem := <-canaldeEntrada2:
				canaldeSaida <- mensagem
			}
		}
	}()

	return canaldeSaida

}

func main() {
	canal := multiplexar(escrever("Olá mundo"), escrever("Programando em Go"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}
