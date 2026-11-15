package ops

type tipo_valor int

const (
	tipo_null tipo_valor = iota
	tipo_numero
	tipo_string
)

type Valor struct {
	tipo  tipo_valor
	valor any
}

func Novo_valor_string(s string) Valor {
	return Valor{
		tipo:  tipo_string,
		valor: s,
	}
}

func Novo_valor_numero(n float64) Valor {
	return Valor{
		tipo:  tipo_numero,
		valor: n,
	}
}
