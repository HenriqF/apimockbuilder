package ops

import (
	"errors"
)

//PREPARAR OPERACAO

func preparar_operacao(input string) ([]token, error) {
	if input == "" {
		return nil, errors.New("operacao vazia")
	}

	toks, err := get_tokens(input)
	if err != nil {
		return nil, err
	}

	_, err = ordenar_tokens(toks)
	if err != nil {
		return nil, err
	}

	return toks, nil
}

func Executar(input string) (int, error) {
	_, err := preparar_operacao(input)
	if err != nil {
		return 0, err
	}

	return 67, nil
}
