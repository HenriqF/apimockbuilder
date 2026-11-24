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
						["content-type", "application/json; charset=utf-8"],
						["Set-Cookie", "coisa=233"],
						["Set-Cookie", "tema=escuro"]
					]
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
