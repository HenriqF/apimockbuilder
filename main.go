package main

import (
	ops "apimockbuilder/operacoes"
	"fmt"
)

func main() {

	fmt.Println(ops.Executar("  'its amazing' == 'coisa legal\\'  "))
}
