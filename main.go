package main

import (
	ops "apimockbuilder/operacoes"
	"fmt"
)

func main() {
	fmt.Println(ops.Executar("maiusculo 'coisa' == corte minusculo '  COISA '"))
}
