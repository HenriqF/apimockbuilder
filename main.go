package main

import (
	"apimockbuilder/api"
	ops "apimockbuilder/operacoes"
	"fmt"
)

func main() {
	operacao := `[
		["evento"],
		["evento"],
		["resposta"],
		["evento"],
		["evento"]
	]`

	ambiente := make(map[string]ops.Valor)
	ambiente_r, resposta, err := api.Executar_regras(operacao, ambiente)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("final:", ambiente_r)
	fmt.Println("resposta:", resposta.Status)
}
