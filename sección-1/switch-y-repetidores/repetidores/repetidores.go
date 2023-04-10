package main

import "fmt"

type Planeta struct {
	name   string
	moonsq int
}

func mostrarPlanetas(planetas []Planeta) {
	for i := 0; i < len(planetas); i++ {

		fmt.Printf("\nEl planeta %v tiene %v lunas.", planetas[i].name, planetas[i].moonsq)

	}
}

func main() {

	mercurio := Planeta{name: "mercurio", moonsq: 0}
	venus := Planeta{name: "venus", moonsq: 0}
	tierra := Planeta{name: "tierra", moonsq: 1}
	marte := Planeta{name: "marte", moonsq: 2}
	jupiter := Planeta{name: "jupiter", moonsq: 63}
	saturno := Planeta{name: "saturno", moonsq: 62}
	urano := Planeta{name: "urano", moonsq: 27}
	neptuno := Planeta{name: "neptuno", moonsq: 13}

	planetas := []Planeta{mercurio, venus, tierra, marte, jupiter, saturno, urano, neptuno}

	if len(planetas) != 0 {
		mostrarPlanetas(planetas)
	} else {
		fmt.Print("No se encontraron planetas")
	}

}
