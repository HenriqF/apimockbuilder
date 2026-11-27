package main

import (
	"apimockbuilder/api"
)

func main() {

	api.Start()

	// operacao := `[
	// 	["resposta",
	// 		{
	// 			"status": "200",
	// 			"headers": [
	// 				["@ minusculo maiusculo 'resposta'", "@2+2"]
	// 			],
	// 			"body": {
	// 				"teste": "coisa",
	// 				"resposta": "sim!",
	// 				"expressao": "@420*2"
	// 			}
	// 		}
	// 	]
	// ]`

	// ambiente := make(map[string]ops.Valor)
	// _, resposta, err := api.Executar_regras(operacao, ambiente)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("\n\nresposta:", resposta)
}
