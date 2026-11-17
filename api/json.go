package api

import (
	ops "apimockbuilder/operacoes"
	"encoding/json"
	"errors"
	"fmt"
)

func executar_bloco(bloco []any, ambiente map[string]ops.Valor) (map[string]ops.Valor, error) {
	exp := bloco[0].(string)

	if len(exp) < 1 {
		return ambiente, nil
	}

	switch {
	case exp[0] == '@':
		res, amb_ret, err := ops.Executar(exp[1:], ambiente)
		if err != nil {
			return ambiente, nil
		}
		fmt.Println(res, err)
		return amb_ret, nil
	}

	return ambiente, nil
}

func Executar_regras(input string, ambiente map[string]ops.Valor) error {
	var processado any

	json.Unmarshal([]byte(input), &processado)
	switch processado.(type) {
	case []any:
	default:
		return errors.New("FORMATO NAO SUPORTADO!")
	}

	for _, prox := range processado.([]any) {
		bloco, _ := prox.([]any)
		if len(bloco) < 1 {
			continue
		}
		if _, ok := bloco[0].(string); !ok {
			continue
		}

		amb, err := executar_bloco(bloco, ambiente)
		if err != nil {
			return err
		}
		ambiente = amb
	}

	return nil
}
