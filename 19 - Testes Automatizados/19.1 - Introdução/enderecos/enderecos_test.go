package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEndereco(t *testing.T) {
	//t.Parallel() <- indica que o teste pode ser rodado em paralelo com outros testes

	cenariosDeTeste := []cenarioDeTeste{
		{"Rua ABC", "Rua"},
		{"Avenida Sumaré", "Avenida"},
		{"Praça da Arvore", "Tipo Inválido!"},
		{"Rodovia Anchieta", "Rodovia"},
		{"Estrada de itaquera", "Estrada"},
		{"RUA DOS IMARÉS", "Rua"},
		{"Beco do batman", "Tipo Inválido!"},
		{"AVENIDA SATÉLITE", "Avenida"},
		{"", "Tipo Inválido!"},
	}

	for _, cenario := range cenariosDeTeste {
		retornoRecebido := TipoDeEndereco(cenario.enderecoInserido)

		if retornoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo de endereço recebido %s é diferente do esperado %s",
				retornoRecebido,
				cenario.retornoEsperado,
			)
		}
	}
}
