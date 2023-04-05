package main

import (
	"fmt"
)

func main() {

	maria := Persona{nombre: "Maria", edad: 25, ciudad: "Tegucigalpa", telefono: 99999999}
	carlos := Persona{nombre: "Carlos", edad: 40, ciudad: "Bogota", telefono: 8888855555}

	fmt.Printf("Persona 1: %v\nPersona 2: %v\n", maria, carlos)

	cambiarCiudad(&maria)

	fmt.Printf("Persona 1 con la ciudad actualizada: %v\nPersona 2: %v\n", maria, carlos)

}

func cambiarCiudad(persona *Persona) {

	persona.ciudad = "Valencia"

}

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono int
}
