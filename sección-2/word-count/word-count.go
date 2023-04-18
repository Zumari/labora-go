package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(texto string) map[string]int {
	contador := strings.Fields(texto)
	contadorPalabras := make(map[string]int)

	for _, palabras := range contador {

		contador, ok := contadorPalabras[palabras]
		if ok {
			contador = contador + 1
		}
		if !ok {
			contador = 1
		}
		contadorPalabras[palabras] = contador
	}
	fmt.Println(contadorPalabras)
	return contadorPalabras
}

func main() {

	texto := "I ate a donut. Then I ate another donut. "
	WordCount(texto)
	wc.Test(WordCount)
}
