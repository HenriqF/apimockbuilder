package ops

import (
	"errors"
	"fmt"
)

// coletar e organizar tokens para avaliacao
func preparar_operacao(input string) ([]token, error) {
	if input == "" {
		return nil, errors.New("operacao vazia")
	}

	toks, err := get_tokens(input)
	if err != nil {
		return nil, err
	}

	toks, err = ordenar_tokens(toks)
	if err != nil {
		return nil, err
	}

	return toks, nil
}

func Executar(input string) (int, error) {
	toks, err := preparar_operacao(input)
	if err != nil {
		return 0, err
	}
	//substituicao de nomes

	//eval

	fmt.Println(toks)
	return 67, nil
}
