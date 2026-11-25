package main

import (
	"apimockbuilder/api"
	ops "apimockbuilder/operacoes"
	"fmt"
)

func main() {
	operacao := `[
		["resposta",
			{
				"status": "200",
				"headers": [
					["@ minusculo maiusculo 'resposta'", "@2+2"]
				]
			}
		]
	]`

	ambiente := make(map[string]ops.Valor)
	_, resposta, err := api.Executar_regras(operacao, ambiente)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\n\nresposta:", resposta)
}
