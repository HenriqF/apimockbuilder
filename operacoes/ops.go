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
		return Valor{tipo_numero, res.valor.(float64)}, ambiente, nil
	}
	if res.tipo == token_string {
		return Valor{tipo_string, res.valor.(string)}, ambiente, nil
	}

	return Valor{}, ambiente, nil
}
