package main

import (
	"fmt"

	"github.com/ferortiz327/godesde0/variables"
)

func main() {
	estado, texto := variables.ConversionTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
