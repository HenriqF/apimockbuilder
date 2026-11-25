package ops

import "strconv"

type Tipo_valor int

const (
	Tipo_nil Tipo_valor = iota
	Tipo_numero
	Tipo_string
)

type Valor struct {
	Valor any
	Tipo  Tipo_valor
}

func Novo_valor_string(s string) Valor {
	return Valor{
		Tipo:  Tipo_string,
		Valor: s,
	}
}

func Novo_valor_numero(n float64) Valor {
	return Valor{
		Tipo:  Tipo_numero,
		Valor: n,
	}
}

func Novo_valor_nil() Valor {
	return Valor{
		Tipo:  Tipo_nil,
		Valor: nil,
	}
}

func (valor *Valor) Valor_para_string() string {
	switch valor.Tipo {
	case Tipo_numero:
		valor.Tipo = Tipo_string
		valor.Valor = strconv.FormatFloat(valor.Valor.(float64), 'g', -1, 64)
	case Tipo_nil:
		valor.Tipo = Tipo_string
		valor.Valor = "nil"
	}
	return valor.Valor.(string)
}
