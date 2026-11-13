package projetos

import "errors"

var projetos = make(map[string]int)

func Lista_projetos() []string {
	resultado := []string{}

	for proj := range projetos {
		resultado = append(resultado, proj)
	}

	return resultado
}

func Novo_projeto(nome string) error {
	if _, existe := projetos[nome]; existe {
		return errors.New("projeto já exiset")
	}
	projetos[nome] = 1
	return nil
}
