package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type usuario struct {
	Nome   string  `json:"nome"`
	Idade  uint    `json:"idade"`
	Cursos []curso `json:"cursos"`
}

type curso struct {
	//Nome        string `json:"-" -> Ao colocar "-" no campo do json ele não irá aparecer no json convertido`
	Nome        string `json:"nome"`
	Instituicao string `json:"instituicao"`
	Campus      string `json:"campus"`
}

func main() {
	usuarioEmJson := `{
		"nome": "Jacksom Guilherme",
		"idade": 26,
		"cursos": [
			{
				"nome": "Análise e desenvolvimento de sistemas",
				"instituicao": "Faculdade de Tecnologia",
				"campus": "Zona Leste"
			}
		]
	}`

	var u usuario

	if erro := json.Unmarshal([]byte(usuarioEmJson), &u); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(u)

	usuario2EmJson := `{
		"nome": "Lino",
		"idade": "1"
	}`

	u2 := make(map[string]string)

	if erro := json.Unmarshal([]byte(usuario2EmJson), &u2); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(u2)
}
