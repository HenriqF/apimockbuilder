package main

import (
	ops "apimockbuilder/operacoes"
	"fmt"
)

func main() {
	ambiente := map[string]ops.Valor{
		"variavel": ops.Novo_valor_string("abcdef"),
	}

	_, ambiente, err := ops.Executar("2 + 3", ambiente)
	if err != nil {
		fmt.Println(err)
	}
}
