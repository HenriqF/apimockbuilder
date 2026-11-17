package main

import (
	"apimockbuilder/api"
	ops "apimockbuilder/operacoes"
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
				["@nome = 30"]
			] 
		}],

		["@nome > 20", {
			"s" : [
				["AGORA É VERDADAE"]
			],
			"n": [
				["CONTINUA SENDO FALSO!"]
			]
		}]
	]`

	ambiente := make(map[string]ops.Valor)
	api.Executar_regras(operacao, ambiente)
}
