package ops

func Executar(input string, ambiente map[string]Valor) (Valor, map[string]Valor, error) {
	toks, err := preparar_operacao(input)
	if err != nil {
		return Valor{}, nil, err
	}
	//fmt.Println("tokens ordenados:", toks)

	res, ambiente, err := executar_operacao(toks, ambiente)
	if err != nil {
		return Valor{}, nil, err
	}

	if res.tipo == token_numero {
		return Valor{res.valor.(float64), Tipo_numero}, ambiente, nil
	}
	if res.tipo == token_string {
		return Valor{res.valor.(string), Tipo_string}, ambiente, nil
	}

	return Valor{}, ambiente, nil
}
