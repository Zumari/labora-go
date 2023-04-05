package main

import (
	"fmt"
)

func main() {

	peliculas := []string{
		"Titanic", "Shazam", "Harry Potter", "El diablo viste de Prada", "Superman", "La mujer maravilla",
		"Thor", "Spiderman", "Entrenando a papá", "Educando a mamá"}
	fmt.Println("Impresión de la matriz: ", peliculas)
	fmt.Println("Impresión del segundo elemento de la matriz: ", peliculas[1])
	fmt.Println("Impresión de la longitud de la matriz: ", len(peliculas))

}
