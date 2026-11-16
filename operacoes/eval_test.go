package ops

import (
	"testing"
)

func TestExecutar_operacao_simples_numeros(t *testing.T) {
	ambiente := map[string]Valor{}

	casos := []struct {
		entrada   string
		resultado any
	}{
		{"10 + 10", float64(20)},
		{"10 + 10 - 20", float64(0)},
		{"-10 + -10", float64(-20)},
		{"10 * 10", float64(100)},
		{"10 / 10", float64(1)},
		{"2 + 2 * 4", float64(10)},
		{"(2 + 2) * 4", float64(16)},
	}

	for i := 0; i < len(casos); i++ {
		res, _, err := Executar(casos[i].entrada, ambiente)
		if err != nil {
			t.Errorf("Deu erro teste... %v", err)
		}
		if res.valor != casos[i].resultado {
			t.Errorf("%v resultou em %v, esperado: %v", casos[i].entrada, res.valor, casos[i].resultado)
		}
	}
}

func TestExecutar_operacao_simples_numeros_e_variaveis(t *testing.T) {
	ambiente := map[string]Valor{
		"numero":      Novo_valor_numero(10),
		"numero_dois": Novo_valor_numero(2),
	}

	casos := []struct {
		entrada   string
		resultado any
	}{
		{"10 + numero", float64(20)},
		{"numero + numero - 20", float64(0)},
		{"-10 + -numero", float64(-20)},
		{"numero * 10", float64(100)},
		{"10 / numero", float64(1)},
		{"2 + 2 * (numero_dois+numero_dois)", float64(10)},
		{"(numero_dois + 2) * 4", float64(16)},
	}

	for i := 0; i < len(casos); i++ {
		res, _, err := Executar(casos[i].entrada, ambiente)
		if err != nil {
			t.Errorf("Deu erro teste... %v", err)
		}
		if res.valor != casos[i].resultado {
			t.Errorf("%v resultou em %v, esperado: %v", casos[i].entrada, res.valor, casos[i].resultado)
		}
	}
}

func TestExecutar_atribuicao_simples(t *testing.T) {
	ambiente := map[string]Valor{
		"numero": Novo_valor_numero(10),
	}

	casos := []struct {
		entrada  string
		esperado any
	}{
		{"numero = 20", float64(20)},
		// {"numero = 30+30", float64(60)},
		// {"numero=67", float64(67)},
		// {"numero =(55+22+11)", float64(88)},
	}

	for i := 0; i < len(casos); i++ {
		_, ambiente_ret, err := Executar(casos[i].entrada, ambiente)
		if err != nil {
			t.Errorf("Deu erro teste... %v", err)
		}
		ambiente = ambiente_ret

		valor, existe := ambiente["numero"]
		if !existe {
			t.Errorf("Variavel 'numero' parou de existir do ambiente...")
		}
		if valor.valor != casos[i].esperado {
			t.Errorf("Variavel tem valor %v, esperado: %v", valor.valor, casos[i].esperado)
		}

	}
}
