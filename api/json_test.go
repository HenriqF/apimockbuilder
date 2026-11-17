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

		ambiente_r, err := Executar_regras(operacao, ambiente)
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
