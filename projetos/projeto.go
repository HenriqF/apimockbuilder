package projetos

import (
	"errors"
	"fmt"
)

var projetos = make(map[string]int)

func Lista_projetos() []string {
	resultado := []string{}

	for proj := range projetos {
		resultado = append(resultado, proj)
	}

	return resultado
}

func Get_projeto(nome string) error {
	_, existe := projetos[nome]
	if !existe {
		return errors.New("projeto nao existe.")
	}
	return nil
}

func Novo_projeto(nome string) error {
	if _, existe := projetos[nome]; existe {
		return errors.New("projeto já exiset")
	}
	projetos[nome] = 1
	return nil
}

func Novo_endpoint(projeto string, endpoint string, regras string) error {
	fmt.Println(projeto, endpoint, regras)
	return nil
}

func Mudar_endpoint(projeto string, endpoint string, regras string) error {
	fmt.Println(projeto, endpoint, regras)
	return nil
}
