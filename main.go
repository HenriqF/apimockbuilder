package main

import (
	"apimockbuilder/api"
	ops "apimockbuilder/operacoes"
	"fmt"
)

func main() {
	operacao := `[
		["@nome = 10"],
		["@nome > 20", {
			"s" : [
				["É VERDADE"],
				["@nome > 45", {
					"s": [["TAMBÉM É VERDADE"]],
					"n": [["AGORA É FALSO"]]
				}]
			],
			"n" : [
				["É FALSO..."],
				["@nome = 30"],
				["@porraloca = 300"]
			]
		}],
	]`

	ambiente := make(map[string]ops.Valor)
	ambiente_r, err := api.Executar_regras(operacao, ambiente)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ambiente_r)
}
