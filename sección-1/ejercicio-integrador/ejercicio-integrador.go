// CONSULTAR COMO DEBEMOS ORDENAR A LAS PERSONAS, POR NOMBRE EDAD O QUE,
// Y QUE SI EL DATO QUE SE INGRESA POR TECLADO ES LA CANTIDAD DE PERSONAS ORDENADAS DE MANERA ASCENDENTE

package main

import (
	"fmt"
)

func main() {

	var nombre string
	var edad int
	var altura int
	var peso int
	var accion int
	var personas []Persona

	textoInicial(&accion, &nombre, &edad, &altura, &peso, &personas)

}

func textoInicial(accion *int, nombre *string, edad *int, altura *int, peso *int, personas *[]Persona) {

	fmt.Println("¿Qué desea hacer? \n1. Crear persona \n2. Ordenar personas \n3. Buscar persona \n4. Ver personas \n5. Salir \n\nIngrese acción a realizar: ")
	fmt.Scan(accion)
	if *accion > 0 && *accion < 6 {
		acciones(accion, *nombre, *edad, *altura, *peso, *personas)
	} else {
		fmt.Println("Dígito no válido")
	}

}

func acciones(accion *int, nombre string, edad int, altura int, peso int, personas []Persona) {

	switch *accion {
	case 1: // Crear persona nueva
		fmt.Println("Ingrese el nombre: ")
		fmt.Scan(&nombre)
		fmt.Println("Ingrese la edad: ")
		fmt.Scan(&edad)
		fmt.Println("Ingrese la altura: ")
		fmt.Scan(&altura)
		fmt.Println("Ingrese el peso: ")
		fmt.Scan(&peso)

		personas = append(personas, crearPersona(&nombre, &edad, &altura, &peso))

		textoInicial(accion, &nombre, &edad, &altura, &peso, &personas)
	case 2: // Ordenar personas
	case 3: // Buscar persona

		if personas != nil {
			fmt.Println("Ingrese el nombre de la persona a buscar: ")
			fmt.Scan(&nombre)
			personaEncontrada := buscarPersonaNombre(personas, nombre)
			fmt.Printf("\nSe encontró a: \nNombre: %v \nEdad: %v \nAltura: %v \nPeso: %v\n\n", personaEncontrada.nombre, personaEncontrada.edad, personaEncontrada.altura, personaEncontrada.peso)
		} else {
			fmt.Printf("No hay personas registradas\n")
		}

		textoInicial(accion, &nombre, &edad, &altura, &peso, &personas)

	case 4: // Ver personas registradas
		if personas == nil {
			fmt.Printf("No se encontraron personas\n")
		} else {
			fmt.Printf("\nPersonas encontradas:\n")
			fmt.Println("==========================")

			for i := 0; i < len(personas); i++ {
				fmt.Printf("%v. \nNombre: %v \nEdad: %v \nAltura: %v \nPeso: %v\n", i+1, personas[i].nombre, personas[i].edad, personas[i].altura, personas[i].peso)
				fmt.Println("--------------------------")
			}
		}
		textoInicial(accion, &nombre, &edad, &altura, &peso, &personas)
	case 5: // Salir
		break
	default:
		fmt.Print("\nOpción no válida\n\n")
		textoInicial(accion, &nombre, &edad, &altura, &peso, &personas)
	}
}

func crearPersona(nombre *string, edad *int, altura *int, peso *int) Persona {

	persona := Persona{nombre: *nombre, edad: *edad, altura: *altura, peso: *peso}
	return persona

}

func buscarPersonaNombre(personas []Persona, nombreBuscado string) *Persona {

	for i := 0; i < len(personas); i++ {
		if personas[i].nombre == nombreBuscado {
			return &personas[i]
		} else {
			fmt.Printf("No se encontraron personas con ese nombre")
		}
		break
	}

	return new(Persona)

}

type Persona struct {
	nombre string
	edad   int
	altura int
	peso   int
}
