package main

import (
	"fmt"
	"introduca-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Avenida Paulista")

	fmt.Println(tipoEndereco)

	tipoEndereco = enderecos.TipoDeEndereco("Rodovia dos Imigrantes")

	fmt.Println(tipoEndereco)

	tipoEndereco = enderecos.TipoDeEndereco("Rua dos bobos")

	fmt.Println(tipoEndereco)

	tipoEndereco = enderecos.TipoDeEndereco("Praça da Arvore")

	fmt.Println(tipoEndereco)
}
