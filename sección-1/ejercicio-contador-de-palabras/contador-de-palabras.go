package main

import (
	"fmt"
	"strings"
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
	// contador := strings.Fields(texto)
	// contadorPalabras := make(map[string]int)

	// for _, palabras := range contador {

	// 	contador, ok := contadorPalabras[palabras]
	// 	if ok {
	// 		contador = contador + 1
	// 	}
	// 	if !ok {
	// 		contador = 1
	// 	}
	// 	contadorPalabras[palabras] = contador
	// }
	// fmt.Println(contadorPalabras)

	// delimitador := "."
	// textoSeparado := strings.Split(texto, delimitador)

	// var oracionSeparadaPalabras []string
	// contadorPalabras := make(map[string]int)
	// for i := 0; i < len(textoSeparado)-1; i++ {
	// 	oracionSeparadaPalabras = strings.Fields(textoSeparado[i])
	// 	for _, palabrasOracion := range oracionSeparadaPalabras {
	// 		contadorPalabras[palabrasOracion] = i + 1
	// 	}
	// 	fmt.Println(oracionSeparadaPalabras)
	// 	fmt.Println(contadorPalabras)
	// }

	// for i, palabras := range textoSeparado {
	// 	palabras = strings.Fields(textoSeparado[i])
	// 	fmt.Println(i, palabras)
	// }

	// fmt.Print(textoSeparado[0])
	// fmt.Print(textoSeparado[1])

	// palabras := strings.Fields()
	// cantidadPalabras := len(palabras)

	fmt.Println()
}
