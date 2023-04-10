package main

import (
	"fmt"
)

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono int
}

// Implementación del constructor NewPersona
func newPersona(nombre string, edad int, ciudad string, telefono int) Persona {
	return Persona{nombre: nombre, edad: edad, ciudad: ciudad, telefono: telefono}
}

// Implementación del método cambiarCiudad
func (p *Persona) cambiarCiudad(nuevaCiudad string) {
	p.ciudad = nuevaCiudad
}

// Implementación de la función String para la estructura Persona
func (p Persona) String() string {
	return fmt.Sprintf("{nombre: %s, edad: %d, ciudad: %s, telefono: %d}", p.nombre, p.edad, p.ciudad, p.telefono)
}

func main() {

	maria := newPersona("Maria", 25, "Tegucigalpa", 99999999)
	carlos := newPersona("Carlos", 40, "Bogota", 8888855555)

	fmt.Printf("Persona 1: %v\nPersona 2: %v\n", maria, carlos)

	maria.cambiarCiudad("Cartagena")

	fmt.Printf("Persona 1 con la ciudad actualizada: %v\nPersona 2: %v\n", maria, carlos)

}
