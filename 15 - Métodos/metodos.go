package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando dados do usuario %s no banco de dados\n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuario1", 16}
	usuario1.salvar()
	fmt.Println(usuario1.maiorDeIdade())

	usuario2 := usuario{"Jack", 26}
	usuario2.salvar()
	fmt.Println(usuario2.maiorDeIdade())
	fmt.Println(usuario2.idade)

	usuario2.fazerAniversario()

	fmt.Println(usuario2.idade)
}
