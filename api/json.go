package api

import (
	ops "apimockbuilder/operacoes"
	"encoding/json"
	"errors"
	"fmt"
)

type t_ambiente map[string]ops.Valor
type t_casos map[string]any

func executar_acao(acao string) {
	fmt.Println("ACAO:", acao)
}

func executar_expressao(exp string, casos t_casos, ambiente t_ambiente) (t_ambiente, error) {
	res, ambiente_r, err := ops.Executar(exp, ambiente)
	if err != nil {
		return ambiente, err
	}
	ambiente = ambiente_r

	caso := "s"
	if res.Tipo == ops.Tipo_numero && res.Valor == float64(0) {
		caso = "n"
	} else if res.Tipo == ops.Tipo_string {
		caso = res.Valor.(string)
	}

	sub_blocos, existe := casos[caso]
	if !existe {
		return ambiente, nil
	}

	ambiente_r, err = processar_blocos(sub_blocos, ambiente)
	if err != nil{
		return ambiente, err
	}
	ambiente = ambiente_r

	return ambiente, nil
}

func executar_bloco(bloco []any, ambiente t_ambiente) (t_ambiente, error) {
	exp := bloco[0].(string)

	switch {
	case len(exp) < 1:
		return ambiente, nil

	case exp[0] == '@':
		var casos t_casos = nil
		if len(bloco) > 1 {
			casos, _ = bloco[1].(map[string]any)
		}

		ambiente_r, err := executar_expressao(exp[1:], casos, ambiente)
		if err != nil {
			return ambiente, err
		}
		return ambiente_r, nil

	default:
		executar_acao(exp)
		return ambiente, nil
	}
}

func get_check_bloco(candidato any) []any {
	bloco, ok := candidato.([]any)
	if !ok {
		return nil
	}
	if len(bloco) < 1 {
		return nil
	}
	if _, ok := bloco[0].(string); !ok {
		return nil
	}
	return bloco
}

func processar_blocos(blocos any, ambiente t_ambiente) (t_ambiente, error) {
	if _, ok := blocos.([]any); !ok {
		return ambiente, errors.New("Regras malformadas")
	}

	for _, prox := range blocos.([]any) {
		bloco := get_check_bloco(prox)
		if bloco == nil {
			continue
		}

		ambiente_r, err := executar_bloco(bloco, ambiente)
		if err != nil {
			return ambiente, err
		}

		ambiente = ambiente_r
	}

	return ambiente, nil
}

func Executar_regras(input string, ambiente t_ambiente) error {
	var blocos any
	err := json.Unmarshal([]byte(input), &blocos)
	if err != nil {
		return err
	}

	_, err = processar_blocos(blocos, ambiente)
	if err != nil {
		return err
	}

	return nil
}
