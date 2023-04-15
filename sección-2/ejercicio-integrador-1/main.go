package main

import (
	"fmt"
	"integrador/pp2"
	"integrador/pp4"
)

func main() {

	var nombre string
	var edad int
	var altura float64
	var peso float64
	var personas []pp2.Persona
	var accion, criterio int
	i := 0
	for i != 1 {

		fmt.Printf("Ingrese el proceso que desea hacer: \n1. Agregar persona\n2. Ordenar personas \n3. Buscar persona \n4. Calcular imc \n5. Cerrar\n")

		fmt.Scan(&accion)

		switch accion {
		case 1:
			pp2.IngresarPersona(nombre, edad, altura, peso, &personas)
			fmt.Printf("\nGrupo de personas: \n%v\n", personas)
		case 2: // Ordenar personas
			fmt.Println("Criterios desea ordenar a las personas: \n1. Nombre \n2. Edad \n3. Altura \n4. Peso")
			fmt.Println("Escriba el número de criterio: ")
			fmt.Scan(&criterio)
			fmt.Printf("\nPersonas ordenadas: \n%v\n", pp2.OrdenarPersonas(personas, criterio))

		case 3: // Buscar personas
			fmt.Println("¿Mediante cúal criterio desea buscar a las personas? \n1. Nombre \n2. Edad \n3. Altura \n4. Peso")
			fmt.Scan(&criterio)
			fmt.Println("Ingrese el atributo a buscar: ")
			switch criterio {
			case 1:
				var valor string
				fmt.Scan(&valor)
				pp2.BuscarPersona(personas, criterio, valor)
			case 2:
				var valor int
				fmt.Scan(&valor)
				pp2.BuscarPersona(personas, criterio, valor)
			case 3:
				var valor float64
				fmt.Scan(&valor)
				pp2.BuscarPersona(personas, criterio, valor)
			case 4:
				var valor float64
				fmt.Scan(&valor)
				pp2.BuscarPersona(personas, criterio, valor)
			}
		case 4:
			var personaIMC int
			fmt.Printf("Ingrese el indice de la persona a la que le desea calcular el IMC \n%v\n", personas)
			fmt.Scan(&personaIMC)
			if personaIMC > len(personas) {
				fmt.Println("Indice inválido,intentelo de nuevo")
			} else {
				pp2.CalcularIMC(personas[personaIMC])
			}
		case 5:
			i = 1
		default:
			fmt.Println("Opción inválida")
		}
	}

	defer func() {
		pp4.CalcularPesoProm(&personas)
		pp4.MinyMax(personas)
	}()
}
