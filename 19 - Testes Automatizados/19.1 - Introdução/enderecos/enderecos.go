package enderecos

import "strings"

// TipoDeEndereco verifica se um endereço tem um tipo válido e o retorna
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoEmMinuscula := strings.ToLower(endereco)
	primeiraPalavraDoEndereco := strings.Split(enderecoEmMinuscula, " ")[0]

	enderecoTemUmTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return strings.Title(primeiraPalavraDoEndereco)
	}

	return "Tipo Inválido!"
}
