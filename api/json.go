package api

import (
	ops "apimockbuilder/operacoes"
	"encoding/json"
	"errors"
	"strconv"
)

type Resposta struct {
	Status  float64
	Headers map[string][]string

	valida bool
}

type t_ambiente map[string]ops.Valor
type t_casos map[string]any

func executar_resposta(bloco []any, ambiente t_ambiente) (t_ambiente, Resposta, error) {
	if len(bloco) == 1 {
		return ambiente, Resposta{Status: 200, valida: true}, nil
	}

	resposta := Resposta{}
	resposta.valida = true

	informacao := bloco[1].(map[string]any)

	status_string, existe := informacao["status"].(string)
	if existe {
		valor, ambiente_r, err := avaliar_expressao(status_string, ambiente)
		if err == nil && valor.Tipo == ops.Tipo_numero {
			ambiente = ambiente_r
			resposta.Status = valor.Valor.(float64)
		}
	}

	headers, existe := informacao["headers"].([]any)
	if existe {
		for _, e := range headers {
			elemento, ok := e.([]any)
			if !ok || len(elemento) != 2 {
				continue
			}
			chave, oka := elemento[0].(string)
			valor, okb := elemento[1].(string)
			if !oka || !okb {
				continue
			}

			chave_resposta, ambiente_r, err := avaliar_expressao(chave, ambiente)
			if err != nil {
				continue
			}
			valor_resposta, ambiente_r, err := avaliar_expressao(valor, ambiente_r)
			if err != nil {
				continue
			}
			chave = chave_resposta.Valor_para_string()
			valor = valor_resposta.Valor_para_string()

			if resposta.Headers == nil {
				resposta.Headers = make(map[string][]string)
			}

			resposta.Headers[chave] = append(resposta.Headers[chave], valor)
		}
	}

	// fmt.Println(resposta)
	return ambiente, resposta, nil
}

func executar_acao(bloco []any, ambiente t_ambiente) (t_ambiente, Resposta, error) {
	acao := bloco[0].(string)

	if acao == "resposta" {
		return executar_resposta(bloco, ambiente)
	}

	return ambiente, Resposta{}, nil
}

// Interpreta a expressao e retorna seu valor. Também verifica se o caractere inicial é @
func avaliar_expressao(exp string, ambiente t_ambiente) (ops.Valor, t_ambiente, error) {
	if len(exp) == 0 {
		return ops.Novo_valor_nil(), ambiente, nil
	}
	if exp[0] != '@' {
		conversao, err := strconv.Atoi(exp)
		if err == nil {
			return ops.Novo_valor_numero(float64(conversao)), ambiente, nil
		}

		return ops.Novo_valor_string(exp), ambiente, nil
	}

	res, ambiente_r, err := ops.Executar(exp[1:], ambiente)
	if err != nil {
		return ops.Valor{}, ambiente, err
	}
	return res, ambiente_r, nil
}

// Executa também recursivamenet blocos de casos. Retorna uma "Resposta"
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
		ambiente_r, resp, err := executar_acao(bloco, ambiente)
		if err != nil {
			return ambiente, Resposta{}, nil
		}
		return ambiente_r, resp, nil
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

		if resp.valida {
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
