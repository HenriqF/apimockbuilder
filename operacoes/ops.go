package ops

import "fmt"

func Executar(input string, ambiente map[string]Valor) (int, map[string]Valor, error) {
	toks, err := preparar_operacao(input)
	if err != nil {
		return 0, nil, err
	}
	fmt.Println("tokens ordenados:", toks)

	res, ambiente, err := executar_operacao(toks, ambiente)
	if err != nil {
		return 0, nil, err
	}

	fmt.Println("resposta:", res)
	return 20, ambiente, nil
}
