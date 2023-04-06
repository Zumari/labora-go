package main

import (
	"fmt"
)

func main() {

	var x int = 1501 // aca hacemos que x valga lo que queremos que valga
	segmentarValorPorRangos(x)

}

func segmentarValorPorRangos(x int) {

	var s1, s2, s3, s4, s5 int

	if x > 0 {
		if x <= 50 {
			s1 = x
			s2 = 0
			s3 = 0
			s4 = 0
			s5 = 0
		} else if x <= 100 {
			s1 = 50
			s2 = x - s1
			s3 = 0
			s4 = 0
			s5 = 0
		} else if x <= 700 {
			s1 = 50
			s2 = 50
			s3 = x - (s1 + s2)
			s4 = 0
			s5 = 0
		} else if x <= 1500 {
			s1 = 50
			s2 = 50
			s3 = 600
			s4 = x - (s1 + s2 + s3)
			s5 = 0
		} else if x > 1500 {
			s1 = 50
			s2 = 50
			s3 = 600
			s4 = 800
			s5 = x - (s1 + s2 + s3 + s4)
		}
	} else {
		fmt.Println("El valor de x debe ser mayor a 0")
	}

	fmt.Println(s1, s2, s3, s4, s5) // por ahora pueden imprimir en pantalla... más adelante veremos que diferencia hay entre hacer esto y devovler como retorno de la función a estas 5 variables!
}
