package api

import (
	ops "apimockbuilder/operacoes"
	"encoding/json"
	"errors"
	"fmt"
)

type Resposta struct {
	Status int
}

type t_ambiente map[string]ops.Valor
type t_casos map[string]any

func executar_acao(acao string) (Resposta, error) {
	fmt.Println("ACAO:", acao)
	if acao == "resposta" {
		return Resposta{Status: 200}, nil
	}

	return Resposta{}, nil
}

func executar_expressao(exp string, casos t_casos, ambiente t_ambiente) (t_ambiente, Resposta, error) {
	res, ambiente_r, err := ops.Executar(exp, ambiente)
	if err != nil {
		return ambiente, Resposta{}, err
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
		return ambiente, Resposta{}, nil
	}

	ambiente_r, resp, err := processar_blocos(sub_blocos, ambiente)
	if err != nil {
		return ambiente, Resposta{}, err
	}
	ambiente = ambiente_r

	return ambiente, resp, nil
}

func executar_bloco(bloco []any, ambiente t_ambiente) (t_ambiente, Resposta, error) {
	exp := bloco[0].(string)

	switch {
	case len(exp) < 1:
		return ambiente, Resposta{}, nil

	case exp[0] == '@':
		var casos t_casos = nil
		if len(bloco) > 1 {
			casos, _ = bloco[1].(map[string]any)
		}

		ambiente_r, resp, err := executar_expressao(exp[1:], casos, ambiente)
		if err != nil {
			return ambiente, Resposta{}, err
		}
		return ambiente_r, resp, nil

	default:
		resp, err := executar_acao(exp)
		if err != nil {
			return ambiente, Resposta{}, nil
		}
		return ambiente, resp, nil
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

func processar_blocos(blocos any, ambiente t_ambiente) (t_ambiente, Resposta, error) {
	if _, ok := blocos.([]any); !ok {
		return ambiente, Resposta{}, errors.New("Regras malformadas")
	}


	for _, prox := range blocos.([]any) {
		bloco := get_check_bloco(prox)
		if bloco == nil {
			continue
		}

		ambiente_r, resp, err := executar_bloco(bloco, ambiente)
		if err != nil {
			return ambiente, Resposta{}, err
		}
		ambiente = ambiente_r

		if resp != (Resposta{}) {
			return ambiente, resp, nil
		}

	}

	return ambiente, Resposta{}, nil
}

func Executar_regras(input string, ambiente t_ambiente) (t_ambiente, Resposta, error) {
	var blocos any
	err := json.Unmarshal([]byte(input), &blocos)
	if err != nil {
		return ambiente, Resposta{}, err
	}

	ambiente_r, resp, err := processar_blocos(blocos, ambiente)
	if err != nil {
		return ambiente, Resposta{}, err
	}

	return ambiente_r, resp, nil
}
