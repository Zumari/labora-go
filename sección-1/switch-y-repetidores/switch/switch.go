package main

import (
	"fmt"
)

func diaSemana(x int) {

	switch x {
	case 1:
		fmt.Println("Lunes")
	case 2:
		fmt.Println("Martes")
	case 3:
		fmt.Println("Miércoles")
	case 4:
		fmt.Println("Jueves")
	case 5:
		fmt.Println("Viernes")
	case 6:
		fmt.Println("Sábado")
	case 7:
		fmt.Println("Domingo")

	}

}

func main() {

	var dia int

	fmt.Println("Ingrese un número del 1 al 7: ")
	fmt.Scan(&dia)
	if dia > 1 && dia < 7 {
		diaSemana(dia)
	} else {
		fmt.Println("Dígito no válido")
	}

}
