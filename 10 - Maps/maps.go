package main

import "fmt"

func main() {
	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Jacksom",
			"ultimo":   "Santos",
		},
		"endereco": {
			"logradouro": "Rua Número Zero",
			"numero":     "99",
			"bairro":     "Itaquera",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome")
	fmt.Println(usuario2)

	usuario2["curso"] = map[string]string{
		"nome":   "Analise de sistemas",
		"campus": "Fatec ZL",
	}

	fmt.Println(usuario2)
}
