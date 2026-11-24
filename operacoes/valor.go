package ops

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
