package ops

import (
	"errors"
)

func bool_to_int(b bool) int {
	if b {
		return 1
	}
	return 0
}

func int_to_bool(i int) bool {
	if i == 0 {
		return false
	}
	return true
}

func get_token_precedence(t token) int {
	if t.valor == "||" {
		return -2
	}
	if t.valor == "&&" {
		return -1
	}
	if t.valor == "==" || t.valor == ">" || t.valor == "<" || t.valor == ">=" || t.valor == "<=" || t.valor == "!=" {
		return 0
	}
	if t.valor == "+" || t.valor == "-" {
		return 1
	}
	if t.valor == "*" || t.valor == "/" {
		return 2
	}
	if t.valor == "u-" {
		return 3
	}
	if t.valor == "!" {
		return 4
	}

	return -67
}

func find_string_end(op string, i int) int {
	pair := rune(op[i])
	size := len(op) - 1

	i++
	for {
		c := rune(op[i])
		ct := get_token_tipo(c)

		if ct == token_delim && c == pair {
			break
		}

		if i == size {
			return -1
		}

		i++
	}

	return i
}

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
