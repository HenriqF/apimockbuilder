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
				"headers": {
							"content-type": "application/json; charset=utf-8",
							"x-verdict": "FALSO!!!",
							"cache-control": "no-store"
						},
				"body": {
							"valido": false,
							"motivo": "FALSO!!!",
							"detalhes": "A afirmação foi rejeitada pela banca do café."
						}
			}
		]
	]`

	ambiente := make(map[string]ops.Valor)
	ambiente_r, resposta, err := api.Executar_regras(operacao, ambiente)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("final:", ambiente_r)
	fmt.Println("resposta:", resposta.Status)
}
