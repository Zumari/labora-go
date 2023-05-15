package main

import (
	"fmt"

	rotarcadenas "github.com/Zumari/labora-go/retos/ADN/rotarCadenas"
)

func main() {
	s := "ABCD"
	a := "DABC"

	fmt.Println(len(s))
	fmt.Println("Cadena original: ", rotarcadenas.PrintCadena(s))
	fmt.Println("Cadena rotada a la derecha: ", rotarcadenas.RotarDerecha(s))
	fmt.Println("Cadena rotada a la izquierda: ", rotarcadenas.RotarIzquierda(s))

	if rotarcadenas.CompararADN(s, a) {
		fmt.Println("Las cadenas de ADN son iguales")
	} else {
		fmt.Println("Las cadenas de ADN son diferentes")
	}

}
