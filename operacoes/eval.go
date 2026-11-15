package ops

import (
	"fmt"
)

var stack []token

func stack_pop() token {
	val := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return val
}
func stack_add(t token) {
	stack = append(stack, t)
}
func pop_bin_op() (token, token) {
	if len(stack) < 2 {
		panic("operacao malformada")
	}
	b, a := stack_pop(), stack_pop()
	return a, b
}
func pop_un_op() token {
	if len(stack) < 1 {
		panic("operacao malformada")
	}
	a := stack_pop()
	return a
}

// OPERACAOES
func op_somar() {
	a, b := pop_bin_op()

	switch {
	case a.tipo == token_numero && b.tipo == token_numero:
		stack_add(token{a.valor.(float64) + b.valor.(float64), token_numero})
		return
	}
	panic("operacao malformada")
}
func op_subtrair() {
	a, b := pop_bin_op()

	switch {
	case a.tipo == token_numero && b.tipo == token_numero:
		stack_add(token{a.valor.(float64) - b.valor.(float64), token_numero})
		return
	}
	panic("operacao malformada")
}
func op_mult() {
	a, b := pop_bin_op()

	switch {
	case a.tipo == token_numero && b.tipo == token_numero:
		stack_add(token{a.valor.(float64) * b.valor.(float64), token_numero})
		return
	}
	panic("operacao malformada")
}
func op_div() {
	a, b := pop_bin_op()

	switch {
	case a.tipo == token_numero && b.tipo == token_numero:
		stack_add(token{a.valor.(float64) / b.valor.(float64), token_numero})
		return
	}
	panic("operacao malformada")
}
func op_umenos() {
	a := pop_un_op()
	if a.tipo != token_numero {
		panic("operacao malformada")
	}
	stack_add(token{a.valor.(float64) * -1, token_numero})
}

// --
var operacoes = map[string]func(){
	"+":  op_somar,
	"-":  op_subtrair,
	"*":  op_mult,
	"/":  op_div,
	"u-": op_umenos,
}

//----------

func executar_operacao(tokens []token, ambiente map[string]Valor) (ret_token token, ret_ambiente map[string]Valor, ret_err error) {
	defer func() {
		if err := recover(); err != nil {
			ret_token = token{}
			ret_ambiente = ambiente
			ret_err = fmt.Errorf("%v", err)
		}
	}()
	stack = stack[:0]

	for _, t := range tokens {
		fmt.Println(stack)
		if t.tipo != token_op {
			stack = append(stack, t)
			continue
		}
		if len(stack) == 0 {
			panic("operacao malformada")
		}

		operacao, existe := operacoes[t.valor.(string)]
		if !existe {
			panic("operador desconhecido")
		}
		operacao()
	}

	return stack[0], ambiente, nil
}
