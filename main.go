package main

import (
	ops "apimockbuilder/operacoes"
	"fmt"
)

func main() {
	fmt.Println(ops.Executar("22 == !-44"))
}
