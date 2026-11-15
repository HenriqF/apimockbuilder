package ops

import (
	"errors"
	"fmt"
)

var stack []token

func stack_pop() token {
	val := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return val
}
func stack_add(t token) error {
	stack = append(stack, t)
	return nil
}

func op_somar() error {
	if len(stack) < 2 {
		return errors.New("operacao malformada")
	}
	b := stack_pop()
	a := stack_pop()

	switch {
	case a.tipo == token_numero && b.tipo == token_numero:
		return stack_add(token{a.valor.(float64) + b.valor.(float64), token_numero})
	}
	return errors.New("operacao malformada")
}

var operacoes = map[string]func() error{
	"+": op_somar,
}

func executar_operacao(tokens []token, ambiente map[string]Valor) (token, map[string]Valor, error) {
	stack = stack[:0]

	for _, t := range tokens {
		fmt.Println(stack)
		if t.tipo != token_op {
			stack = append(stack, t)
			continue
		}
		if len(stack) == 0 {
			return token{}, nil, errors.New("operacao malformada")
		}

		operacao, existe := operacoes[t.valor.(string)]
		if !existe {
			return token{}, nil, errors.New("operador desconhecido")
		}

		err := operacao()
		if err != nil {
			return token{}, nil, err
		}
	}

	return stack[0], ambiente, nil
}
