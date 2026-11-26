package ops

import (
	"fmt"
	"strings"
)

var ambiente map[string]Valor

// tenta retornar valor identificado por nome no ambiente.
// cria novo valor com tipo nulo se nome nao existir
func get_ambiente(nome string) token {
	valor, existe := ambiente[nome]
	if !existe {
		novo := Valor{nil, Tipo_nil}
		ambiente[nome] = novo
		valor = novo
	}

	switch valor.Tipo {
	case Tipo_numero:
		return token{valor.Valor, token_numero}
	case Tipo_string:
		return token{valor.Valor, token_string}
	}
	return token{nil, token_nil}
}
func set_ambiente(nome string, valor token) {
	switch valor.tipo {
	case token_numero:
		ambiente[nome] = Valor{valor.valor, Tipo_numero}
		return
	case token_string:
		ambiente[nome] = Valor{valor.valor, Tipo_string}
		return
	}
	ambiente[nome] = Valor{nil, Tipo_nil}
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

// se o topo for um nome, nao buscamos seu valor no ambiente
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

// pop_bin_op so que para operacao de atribuicao (a nao deve ser convertido de nome para valor)
func pop_bin_op_atr() (token, token) {
	b, a := pop_un_op(), pop_un_op_nosolve()
	return a, b
}

// OPERACAOES
func op_somar() {
	a, b := pop_bin_op()

	switch {
	case a.tipo != b.tipo:
		panic("operacao malformada")
	case a.tipo == token_numero:
		stack_add(token{a.valor.(float64) + b.valor.(float64), token_numero})
		return
	case a.tipo == token_string:
		stack_add(token{a.valor.(string) + b.valor.(string), token_string})
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
	case a.tipo == token_string && b.tipo == token_numero:
		resultado := a.valor.(string)
		for i := 1; i < (int)(b.valor.(float64)); i++ {
			resultado = resultado + a.valor.(string)
		}
		stack_add(token{resultado, token_string})
	default:
		panic("operacao malformada")
	}
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
func op_igual() {
	a, b := pop_bin_op()

	switch {
	case a.tipo != b.tipo:
		stack_add(token{float64(0), token_numero})
	case a.valor == b.valor:
		stack_add(token{float64(1), token_numero})
	default:
		stack_add(token{float64(0), token_numero})
	}
}
func op_diff() {
	a, b := pop_bin_op()

	switch {
	case a.tipo != b.tipo:
		stack_add(token{float64(1), token_numero})
	case a.valor != b.valor:
		stack_add(token{float64(1), token_numero})
	default:
		stack_add(token{float64(0), token_numero})
	}
}

func op_comparacao(comparacao func(float64, float64) bool) {
	a, b := pop_bin_op()

	switch {
	case a.tipo != b.tipo:
		panic("comparacao com tipos indevidos")
	case a.tipo == token_string:
		panic("comparacao com strings")
	case comparacao(a.valor.(float64), b.valor.(float64)):
		stack_add(token{float64(1), token_numero})
	default:
		stack_add(token{float64(0), token_numero})
	}
}
func op_maior() {
	op_comparacao(func(a, b float64) bool { return a > b })
}
func op_menor() {
	op_comparacao(func(a, b float64) bool { return a < b })
}
func op_maiorigual() {
	op_comparacao(func(a, b float64) bool { return a >= b })
}
func op_menorigual() {
	op_comparacao(func(a, b float64) bool { return a <= b })
}

func op_maiusculo() {
	a := pop_un_op()
	if a.tipo != token_string {
		panic("operacao malformada")
	}
	stack_add(token{strings.ToUpper(a.valor.(string)), token_string})
}
func op_minusculo() {
	a := pop_un_op()
	if a.tipo != token_string {
		panic("operacao malformada")
	}
	stack_add(token{strings.ToLower(a.valor.(string)), token_string})
}
func op_corte() {
	a := pop_un_op()
	if a.tipo != token_string {
		panic("operacao malformada")
	}
	stack_add(token{strings.TrimSpace(a.valor.(string)), token_string})
}
func op_contem() {
	a, b := pop_bin_op()
	if a.tipo != token_string || b.tipo != token_string {
		panic("operacao malformada")
	}
	if strings.Contains(a.valor.(string), b.valor.(string)) {
		stack_add(token{float64(1), token_numero})
	} else {
		stack_add(token{float64(0), token_numero})
	}
}
func op_inicia() {
	a, b := pop_bin_op()
	if a.tipo != token_string || b.tipo != token_string {
		panic("operacao malformada")
	}
	if strings.HasPrefix(b.valor.(string), a.valor.(string)) {
		stack_add(token{float64(1), token_numero})
	} else {
		stack_add(token{float64(0), token_numero})
	}
}
func op_finaliza() {
	a, b := pop_bin_op()
	if a.tipo != token_string || b.tipo != token_string {
		panic("operacao malformada")
	}
	if strings.HasSuffix(b.valor.(string), a.valor.(string)) {
		stack_add(token{float64(1), token_numero})
	} else {
		stack_add(token{float64(0), token_numero})
	}
}

// --
var operacoes = map[string]func(){
	"+":  op_somar,
	"-":  op_subtrair,
	"*":  op_mult,
	"/":  op_div,
	"u-": op_umenos,
	"=":  op_atribuir,

	"==": op_igual,
	"!=": op_diff,
	">":  op_maior,
	"<":  op_menor,
	">=": op_maiorigual,
	"<=": op_menorigual,

	"maiusculo": op_maiusculo,
	"minusculo": op_minusculo,
	"corte":     op_corte,
	"contem":    op_contem,
	"inicia":    op_inicia,
	"finaliza":  op_finaliza,
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
		// fmt.Println("... ", stack)
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
