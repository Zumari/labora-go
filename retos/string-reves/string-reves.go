package main

import (
	"fmt"
	"strings"
)

func stringVolteado(s *string) string {

	textoSeparado := []rune(*s)
	textoVolteado := make([]rune, len(textoSeparado))
	for i, j := 0, len(textoSeparado)-1; i < len(textoSeparado); i, j = i+1, j-1 {
		textoVolteado[i] = textoSeparado[j]
	}
	return string(textoVolteado)

}

func palindromo(texto *string) {
	textoVolteado := strings.ToLower(stringVolteado(texto))
	minusTexto := strings.ToLower(*texto)
	if minusTexto != textoVolteado {
		fmt.Println("El texto no es palíndromo")
	} else {
		fmt.Println("El texto es palíndromo")
	}

}

func main() {

	var texto string
	fmt.Println("Ingrese el texto que desea escribir al revés")
	fmt.Scan(&texto)
	fmt.Println("El texto escrito al revés es: ", stringVolteado(&texto))
	palindromo(&texto)

}
