package rotarcadenas

import (
	"strings"
)

func PrintCadena(c string) string {
	return c
}

func RotarDerecha(c string) string {
	var cadena, cOrdenada []string
	inicial := string(c[len(c)-1])
	cOrdenada = append(cOrdenada, inicial)

	for _, v := range c {
		cadena = append(cadena, string(v))
	}

	for i := 0; i < len(cadena)-1; i++ {
		cOrdenada = append(cOrdenada, cadena[i])
	}

	return strings.Join(cOrdenada, "")
}

func RotarIzquierda(c string) string {
	var cadena, cOrdenada []string

	for _, v := range c {
		cadena = append(cadena, string(v))
	}

	for i := 0; i < len(cadena)-1; i++ {
		cOrdenada = append(cOrdenada, cadena[i+1])
	}

	cOrdenada = append(cOrdenada, cadena[0])

	return strings.Join(cOrdenada, "")
}

func CompararADN(a, b string) bool {
	rotando := a
	var resultado bool
	for i := 0; i < len(a); i++ {
		rotando = RotarDerecha(rotando)
		if rotando == b {
			resultado = true
		}
	}
	return resultado
}
