package api

import (
	ops "apimockbuilder/operacoes"
	"testing"
)

func Test_leitura_escrita_ambiente(t *testing.T) {
	operacao := `[
		["@numero >= 67", {
			"s": [
				["@numero = 1"]
			],
			"n": [
				["@numero = 0"]
			]
		}]
	]`

	casos := []struct {
		entrada, saida float64
	}{
		{67, 1},
		{68, 1},
		{0, 0},
		{66.999, 0},
	}

	for _, caso := range casos {
		ambiente := make(map[string]ops.Valor)
		ambiente["numero"] = ops.Novo_valor_numero(caso.entrada)

		ambiente_r, _, err := Executar_regras(operacao, ambiente)
		if err != nil {
			t.Errorf("%v", err)
		}

		valor, ok := ambiente_r["numero"]
		if !ok {
			t.Errorf("Variável 'numero' sumiu do ambiente..")
		}
		if valor.Valor != caso.saida {
			t.Errorf("Valor errado retornado: %v, esperado: %v", valor.Valor, caso.saida)
		}
	}
}

func Test_operacao_retorna_com_resposta(t *testing.T) {
	operacao := `[
		["@numero == numero", {
			"s": [
				["resposta"]
			]
		}],
		["@numero = numero + 1"]
	]`

	casos := []float64{1, 5, 20, 23, 44, 123}

	for _, caso := range casos {
		ambiente := make(map[string]ops.Valor)
		ambiente["numero"] = ops.Novo_valor_numero(caso)

		ambiente_r, _, err := Executar_regras(operacao, ambiente)
		if err != nil {
			t.Errorf("%v", err)
		}

		valor, ok := ambiente_r["numero"]
		if !ok {
			t.Errorf("Variável 'numero' sumiu do ambiente..")
		}
		if valor.Valor != caso {
			t.Errorf("Valor errado retornado: %v, esperado: %v", valor.Valor, caso)
		}
	}
}

func Test_operacao_retorna_com_resposta_e_valida_status(t *testing.T) {
	operacao := `[
		["resposta", {
			"status": "@numero"
		}]
	]`
	casos := []float64{1, 5, 6, 67, 1923, 399, 179584, 0, 200, 400, 500, 100}

	for _, caso := range casos {
		ambiente := make(map[string]ops.Valor)
		ambiente["numero"] = ops.Novo_valor_numero(caso)

		_, resposta, err := Executar_regras(operacao, ambiente)
		if err != nil {
			t.Errorf("%v", err)
		}

		valor := resposta.Status
		if valor != caso {
			t.Errorf("Valor errado retornado: %v, esperado: %v", valor, caso)
		}
	}
}

func Test_operacao_retorna_com_resposta_e_valida_headers(t *testing.T) {
	operacao := `[
		["resposta",
			{
				"headers": [
					["@ maiusculo palavra", "@numero+1"]
				]
			}
		]
	]`

	casos := []struct {
		palavra    string
		numero     int
		p_esperado string
		n_esperado string
	}{
		{"coisa", 1, "COISA", "2"},
		{"header", 30, "HEADER", "31"},
		{"ABCDEF", 20, "ABCDEF", "21"},
	}

	for _, caso := range casos {
		ambiente := make(map[string]ops.Valor)
		ambiente["numero"] = ops.Novo_valor_numero(float64(caso.numero))
		ambiente["palavra"] = ops.Novo_valor_string(string(caso.palavra))

		_, resposta, err := Executar_regras(operacao, ambiente)
		if err != nil {
			t.Errorf("%v", err)
		}

		valor, existe := resposta.Headers[string(caso.p_esperado)]
		if !existe {
			t.Errorf("Chave '%v' nao existe dentro dos headers..", caso.p_esperado)
		}

		if len(valor) != 1 {
			t.Errorf("Chave '%v' contem mais valores %v do que o esperado...", caso.p_esperado, valor)
		}

		if valor[0] != caso.n_esperado {
			t.Errorf("Valor de header errado: %v, esperado: %v", valor, caso.n_esperado)
		}
	}
}
