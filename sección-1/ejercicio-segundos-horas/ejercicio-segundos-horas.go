package main

import (
	"fmt"
)

func main() {

	seg := 0
	min := 0
	hor := 0
	dia := 0

	fmt.Println("Escriba la cantidad de segundos a convertir:")
	fmt.Scan(&seg)

	if seg > 0 {

		if seg < 60 {
			fmt.Printf("\n=================================== \nEquivale a: \n%v días. \n%v horas. \n%v minutos. \n%v segundos. \n", dia, hor, min, seg)
		} else if seg >= 60 && seg < 3600 {

			min = seg / 60
			seg = seg % 60

			fmt.Printf("\n=================================== \nEquivale a: \n%v días. \n%v horas. \n%v minutos. \n%v segundos. \n", dia, hor, min, seg)

		} else if seg >= 3600 && seg < 86400 {

			min = seg / 60
			hor = min / 60
			min = min % 60
			seg = seg % 60

			fmt.Printf("\n=================================== \nEquivale a: \n%v días. \n%v horas. \n%v minutos. \n%v segundos. \n", dia, hor, min, seg)

		} else if seg >= 86400 {

			min = seg / 60
			hor = min / 60
			dia = hor / 24
			hor = hor % 24
			min = min % 60
			seg = seg % 60

			fmt.Printf("\n=================================== \nEquivale a: \n%v días. \n%v horas. \n%v minutos. \n%v segundos. \n", dia, hor, min, seg)

		} else {
			fmt.Println("El número debe ser mayor que cero")
		}

	}

}
