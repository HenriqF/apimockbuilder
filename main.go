package main

import (
	"apimockbuilder/api"
	ops "apimockbuilder/operacoes"
	"fmt"
)

func main() {
	operacao := `[
		["@nome = 20"],
		["@nome > 20", {
			"s" : ["resposta", {
						"status": 200,
						"body": {
							"mensagem" : "@nome"
						}
					}], 

			"n" : ["resposta", {
						"status": 200,
						"headers": {
							"Content-Type": "text/plain"
						},
						"body": "FALSO!!!"
					}] 
		}]
	]`

	ambiente := make(map[string]ops.Valor)
	err := api.Executar_regras(operacao, ambiente)
	fmt.Println(err)
}
