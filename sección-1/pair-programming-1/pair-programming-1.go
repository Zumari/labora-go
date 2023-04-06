package main

import (
	"fmt"
)

type Persona struct {
	nombre   string
	edad     int
	ciudad   string
	telefono int64
}

func main() {

	maria := Persona{nombre: "Maria", edad: 25, ciudad: "Tegucigalpa", telefono: 99999999}
	ludmila := Persona{nombre: "Ludmila", edad: 20, ciudad: "Buenos Aires", telefono: 9829345729845}
	carmelo := Persona{nombre: "Carmelo", edad: 24, ciudad: "Tegucigalpa", telefono: 99999999}
	alonso := Persona{nombre: "Alonso", edad: 32, ciudad: "Buenos Aires", telefono: 9829345729845}
	rodrigo := Persona{nombre: "Rodrigo", edad: 29, ciudad: "Buenos Aires", telefono: 9829345729845}
	fmt.Printf("Valores iniciales \nPersona 1: %v\nPersona 2: %v\n", maria, ludmila)

	incrementarEdad(&ludmila)
	fmt.Printf("Valores incrementando edad \nPersona 1: %v\nPersona 2: %v\n", maria, ludmila)
	personas := []Persona{maria, ludmila, carmelo, alonso, rodrigo}
	fmt.Print(personas)
	buscar1 := buscarEdad(personas, 24)
	if buscar1.edad != 0 {
		fmt.Printf("Nombre: %s, Edad: %d\n", buscar1.nombre, buscar1.edad)
	} else {
		fmt.Printf("No hay una persona registradad con esa edad")
	}

	buscar2 := buscarEdad(personas, 40)
	if buscar2.edad != 0 {
		fmt.Printf("Nombre: %s, Edad: %d\n", buscar2.nombre, buscar2.edad)
	} else {
		fmt.Printf("No hay una persona registradad con esa edad")
	}

	nuevaPersona := crearPersona("Mauro", 25, "Pinamar", 123456789)
	fmt.Println("Nueva persona registrada: ", nuevaPersona)

	actualizarTelefono(&nuevaPersona, 987654321)

	fmt.Printf("El telefono de la persona %s ha sido actualizado, el nuevo numero es: %d\n", nuevaPersona.nombre, nuevaPersona.telefono)
}

func incrementarEdad(persona *Persona) {
	persona.edad++
}
func buscarEdad(personas []Persona, edad int) *Persona {
	for i := 0; i < len(personas); i++ {
		if personas[i].edad == edad {
			return &personas[i]
			break
		}
	}
	return new(Persona)
}

func crearPersona(nombre string, edad int, ciudad string, telefono int64) Persona {
	persona := Persona{nombre: nombre, edad: edad, ciudad: ciudad, telefono: telefono}
	return persona

}

func actualizarTelefono(persona *Persona, telefono int64) {
	persona.telefono = telefono
}
