package main

import (
	ops "apimockbuilder/operacoes"
	"fmt"
)

func main() {
	ambiente := map[string]ops.Valor{
		"numero": ops.Novo_valor_numero(67),
		"string": ops.Novo_valor_string("fdbnasjfk"),
	}

	_, ambiente, err := ops.Executar("numero + numero", ambiente)
	if err != nil {
		fmt.Println(err)
	}
}
