package ops

import (
	"fmt"
)

var ambiente map[string]Valor

// tenta retornar valor identificado por nome no ambiente.
// cria novo valor com tipo nulo se nome nao existir
func get_ambiente(nome string) token {
	valor, existe := ambiente[nome]
	if !existe {
		novo := Valor{tipo_nil, nil}
		ambiente[nome] = novo
		valor = novo
	}

	switch valor.tipo {
	case tipo_numero:
		return token{valor.valor, token_numero}
	case tipo_string:
		return token{valor.valor, token_string}
	}
	return token{nil, token_nil}
}
func set_ambiente(nome string, valor token) {
	switch valor.tipo {
	case token_numero:
		ambiente[nome] = Valor{tipo_numero, valor.valor}
		return
	case token_string:
		ambiente[nome] = Valor{tipo_string, valor.valor}
		return
	}
	ambiente[nome] = Valor{tipo_nil, nil}
}

var stack []token

func stack_pop() token {
	val := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return val
}
func stack_add(t token) {
	stack = append(stack, t)
}

func pop_un_op_nosolve() token {
	if len(stack) < 1 {
		panic("operacao malformada")
	}
	return stack_pop()
}
func pop_un_op() token {
	if len(stack) < 1 {
		panic("operacao malformada")
	}
	a := stack_pop()

	if a.tipo == token_nome {
		a = get_ambiente(a.valor.(string))
	}

	return a
}
func pop_bin_op() (token, token) {
	b, a := pop_un_op(), pop_un_op()

	return a, b
}
func pop_bin_op_atr() (token, token) {
	b, a := pop_un_op(), pop_un_op_nosolve()
	return  a, b
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
func op_atribuir() {
	a, b := pop_bin_op_atr()

	if a.tipo != token_nome {
		panic("operacao malformada")
	}

	stack_add(b)

	set_ambiente(a.valor.(string), b)
}

// --
var operacoes = map[string]func(){
	"+":  op_somar,
	"-":  op_subtrair,
	"*":  op_mult,
	"/":  op_div,
	"u-": op_umenos,
	"=":  op_atribuir,
}

//----------

func executar_operacao(tokens []token, amb map[string]Valor) (ret_token token, ret_amb map[string]Valor, ret_err error) {
	defer func() {
		if err := recover(); err != nil {
			ret_token = token{}
			ret_amb = ambiente
			ret_err = fmt.Errorf("%v", err)
		}
	}()
	stack = stack[:0]
	ambiente = amb

	for _, t := range tokens {
		fmt.Println("... ", stack)
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

	if len(stack) < 1 {
		panic("operacao malformada")
	}
	if stack[0].tipo == token_nome {
		stack[0] = get_ambiente(stack[0].valor.(string))
	}

	return stack[0], ambiente, nil
}
