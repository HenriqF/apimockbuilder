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

func Test_operacao_acaba_quando_encontra_acao_resposta(t *testing.T) {
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

func Test_interpretacao_de_status_simples_em_acao_resposta_funciona(t *testing.T) {
	operacao := `[
		["resposta", {
			"status": "200"
		}]
	]`

	ambiente := make(map[string]ops.Valor)
	_, resposta, err := Executar_regras(operacao, ambiente)

	if err != nil {
		t.Errorf("%v", err)
	}

	valor := resposta.Status
	if valor != 200 {
		t.Errorf("Valor errado retornado: %v, esperado: 200", valor)
	}
}

func Test_interpretacao_de_status_em_acao_resposta_funciona(t *testing.T) {
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

func Test_interpretacao_de_headers_em_acao_resposta_funciona(t *testing.T) {
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

func Test_interpretacao_de_body_simples_mapa_em_acao_resposta_funciona(t *testing.T) {
	operacao := `[
		["resposta",
			{
				"body": {
					"chave" : "valor"
				}
			}
		]
	]`

	ambiente := make(map[string]ops.Valor)
	_, resposta, err := Executar_regras(operacao, ambiente)
	if err != nil {
		t.Errorf("%v", err)
	}

	mapa, ok := resposta.Corpo.(map[string]any)
	if !ok {
		t.Errorf("Corpo malformado... %v", resposta.Corpo)
	}
	valor, existe := mapa["chave"]
	if !existe {
		t.Errorf("Chave 'chave' está faltando no corpo...")
	}
	if valor != "valor" {
		t.Errorf("Valor de 'chave' deveria ser 'valor', mas é %v", valor)
	}
}

func Test_interpretacao_de_body_simples_string_em_acao_resposta_funciona(t *testing.T) {
	operacao := `[
		["resposta",
			{
				"body": "valor"
			}
		]
	]`

	ambiente := make(map[string]ops.Valor)
	_, resposta, err := Executar_regras(operacao, ambiente)
	if err != nil {
		t.Errorf("%v", err)
	}

	corpo, ok := resposta.Corpo.(string)
	if !ok {
		t.Errorf("Corpo malformado... %v", resposta.Corpo)
	}
	if corpo != "valor" {
		t.Errorf("Valor de 'chave' deveria ser 'valor', mas é %v", corpo)
	}
}

func Test_interpretacao_de_body_em_acao_resposta_funciona(t *testing.T) {
	operacao := `[
		["resposta",
			{
				"body": {
					"a" : "@ maiusculo palavra",
					"b" : "@ 1 + numero"
				}
			}
		]
	]`

	casos := []struct {
		palavra    string
		numero     float64
		p_esperado string
		n_esperado float64
	}{
		{"coisa", 1, "COISA", 2},
		{"header", 30, "HEADER", 31},
		{"ABCDEF", 20, "ABCDEF", 21},
	}

	for _, caso := range casos {
		ambiente := make(map[string]ops.Valor)
		ambiente["numero"] = ops.Novo_valor_numero(float64(caso.numero))
		ambiente["palavra"] = ops.Novo_valor_string(string(caso.palavra))

		_, resposta, err := Executar_regras(operacao, ambiente)
		if err != nil {
			t.Errorf("%v", err)
		}

		mapa, ok := resposta.Corpo.(map[string]any)
		if !ok {
			t.Errorf("Corpo malformado... %v", resposta.Corpo)
		}

		valor, existe := mapa["a"]
		if !existe {
			t.Errorf("Chave 'a' está faltando no corpo...")
		}
		if valor != caso.p_esperado {
			t.Errorf("Valor de 'a' deveria ser '%v', mas é '%v'", caso.p_esperado, valor)
		}

		valor, existe = mapa["b"]
		if !existe {
			t.Errorf("Chave 'b' está faltando no corpo...")
		}
		if valor != caso.n_esperado {
			t.Errorf("Valor de 'b' deveria ser '%v', mas é '%v'", caso.n_esperado, valor)
		}

	}
}
