package main

import "fmt"

func main() {

	var equipo1 = map[string]string{
		"nombre": "Lobos",
		"valor":  "12",
	}

	var equipo2 = map[string]string{
		"nombre": "Estrellas",
		"valor":  "28",
	}

	var equipo3 = map[string]string{
		"nombre": "Cachorros",
		"valor":  "25",
	}

	var equipo4 = map[string]string{
		"nombre": "Hormigas",
		"valor":  "16",
	}

	var equipo5 = map[string]string{
		"nombre": "Gatos",
		"valor":  "35",
	}

	fmt.Printf("El equipo %v ha ganado %v títulos.", equipo1["nombre"], equipo1["valor"])
	fmt.Printf("\nEl equipo %v ha ganado %v títulos.", equipo2["nombre"], equipo2["valor"])
	fmt.Printf("\nEl equipo %v ha ganado %v títulos.", equipo3["nombre"], equipo3["valor"])
	fmt.Printf("\nEl equipo %v ha ganado %v títulos.", equipo4["nombre"], equipo4["valor"])
	fmt.Printf("\nEl equipo %v ha ganado %v títulos.", equipo5["nombre"], equipo5["valor"])

}
