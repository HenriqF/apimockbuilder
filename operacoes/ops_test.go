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

	for i := range casos {
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

	for i := range casos {
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
		{"numero = 30+30", float64(60)},
		{"numero=67", float64(67)},
		{"numero =(55+22+11)", float64(88)},
		{"numero = 'coisa'", "coisa"},
	}

	for i := range casos {
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
			t.Errorf("%v\nVariavel tem valor %v, esperado: %v", casos[i].entrada, valor.valor, casos[i].esperado)
		}
	}
}

func TestExecutar_atribuicao_extra(t *testing.T) {
	ambiente := map[string]Valor{
		"numero": Novo_valor_numero(10),
	}

	casos := []struct {
		entrada  string
		esperado any
	}{
		{"numero = 20", float64(20)},
		{"numero", float64(20)},
		{"nova = 40", float64(40)},
		{"nova", float64(40)},
		{"nova = 40 + 20", float64(60)},
		{"(nova = 40 + 20) / numero", float64(3)},
	}

	for i := range casos {
		res, ambiente_ret, err := Executar(casos[i].entrada, ambiente)
		if err != nil {
			t.Errorf("Deu erro teste... %v", err)
		}
		ambiente = ambiente_ret

		if res.valor != casos[i].esperado {
			t.Errorf("%v\nExpressão resultou em %v. Esperado: %v", casos[i].entrada, res.valor, casos[i].esperado)
		}
	}
}

func TestExecutar_comparacoes(t *testing.T) {
	ambiente := map[string]Valor{
		"numero": Novo_valor_numero(10),
	}

	casos := []struct {
		entrada  string
		esperado any
	}{
		{"numero == 10", float64(1)},
		{"numero == 20", float64(0)},
		{"numero != 20", float64(1)},
		{"numero != 10", float64(0)},
		{"numero > 5", float64(1)},
		{"numero > 10", float64(0)},
		{"numero < 20", float64(1)},
		{"numero < 10", float64(0)},
		{"numero >= 10", float64(1)},
		{"numero >= 15", float64(0)},
		{"numero <= 10", float64(1)},
		{"numero <= 5", float64(0)},
		{"'def' == 'abc'", float64(0)},
		{"'def' != 'abc'", float64(1)},
		{"'abc' == 'abc'", float64(1)},
		{"'def' == 12", float64(0)},
		{"'def' != 12", float64(1)},
	}

	for i := range casos {
		res, _, err := Executar(casos[i].entrada, ambiente)
		if err != nil {
			t.Errorf("Deu erro teste... %v", err)
		}
		if res.valor != casos[i].esperado {
			t.Errorf("%v resultou em %v, esperado: %v", casos[i].entrada, res.valor, casos[i].esperado)
		}
	}
}

func TestExecutar_operacoes_strings(t *testing.T) {
	ambiente := map[string]Valor{}

	casos := []struct {
		entrada  string
		esperado any
	}{
		{"'teste'", "teste"},
		{"maiusculo 'teste'", "TESTE"},
		{"minusculo 'TESTE'", "teste"},
		{"corte '  TESTE  '", "TESTE"},

		{"'coisa' contem 'isa'", float64(1)},
		{"'coisa' contem 'lisa'", float64(0)},
		{"'co' inicia 'coisa'", float64(1)},
		{"'co' inicia 'cuisa'", float64(0)},
		{"'isa' finaliza 'cuisa'", float64(1)},
		{"'lisa' finaliza 'cuisa'", float64(0)},

		{"'coisa' + ' legal'", "coisa legal"},
		{"'coisa' * 2", "coisacoisa"},
	}

	for i := range casos {
		res, _, err := Executar(casos[i].entrada, ambiente)
		if err != nil {
			t.Errorf("Deu erro teste... %v", err)
		}
		if res.valor != casos[i].esperado {
			t.Errorf("%v resultou em %v, esperado: %v", casos[i].entrada, res.valor, casos[i].esperado)
		}
	}
}

func TestExecutar_mistura(t *testing.T) {
	ambiente := map[string]Valor{
		"a": Novo_valor_numero(10),
		"b": Novo_valor_numero(5),
		"c": Novo_valor_numero(2),
	}

	casos := []struct {
		entrada  string
		esperado any
	}{
		{"(a + b) * c - (a / b)", float64(28)},
		{"(a * b > 25) == 1", float64(1)},
		{"('porraloca' contem 'rra') == 1", float64(1)},
		{"((a + b) * c) != (b * c * a)", float64(1)},
		{"maiusculo ('teste ' + 'foda') + ' - v67'", "TESTE FODA - v67"},
		{"corte '   sigma da bahia   ' + ' (ok)'", "sigma da bahia (ok)"},
		{"minusculo maiusculo 'eeeeeeeee' + ' 6767'", "eeeeeeeee 6767"},
	}

	for i := range casos {
		res, _, err := Executar(casos[i].entrada, ambiente)
		if err != nil {
			t.Errorf("Deu erro teste... %v", err)
		}
		if res.valor != casos[i].esperado {
			t.Errorf("%v resultou em %v, esperado: %v", casos[i].entrada, res.valor, casos[i].esperado)
		}
	}
}
