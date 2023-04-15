package pp2

import (
	"fmt"
	"sort"
)

type Persona struct {
	Nombre string
	Edad   int
	Altura float64
	Peso   float64
}

func CrearPersona(nombre *string, edad *int, altura *float64, peso *float64) *Persona {

	if *edad < 0 || *altura < 0 || *peso < 0 {
		fmt.Println("Error en alguno de los datos ingresados")
		return nil
	}
	return &Persona{Nombre: *nombre, Edad: *edad, Altura: *altura, Peso: *peso}

}

func OrdenarPersonas(personas []Persona, criterio int) []Persona {
	switch criterio {
	case 1:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].Nombre < personas[j].Nombre
		})
	case 2:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].Edad < personas[j].Edad
		})
	case 3:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].Altura < personas[j].Altura
		})
	case 4:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].Peso < personas[j].Peso
		})
	default:
	}

	return personas
}

func BuscarPersona(personas []Persona, criterio int, valor any) {
	var comparar func(p Persona) bool
	switch criterio {
	case 1:
		comparar = func(p Persona) bool {
			return p.Nombre == valor.(string)
		}
	case 2:
		comparar = func(p Persona) bool {
			return p.Edad == valor.(int)
		}
	case 3:
		comparar = func(p Persona) bool {
			return p.Altura == valor.(float64)
		}
	case 4:
		comparar = func(p Persona) bool {
			return p.Peso == valor.(float64)
		}
	default:
		fmt.Println("Criterio de búsqueda no válido.")
		return
	}
	var persona1 Persona
	for i := 0; i < len(personas); i++ {
		if comparar(personas[i]) {
			persona1 = personas[i]
			fmt.Printf("Se encontró a: %v\n", persona1)
			return
		}
	}
	fmt.Print("El atributo no pertenece a ninguna persona\n")
}

func CalcularIMC(persona Persona) {
	IMC := persona.Peso / (persona.Altura * persona.Altura)
	var categoria string

	if IMC < 18.5 {
		categoria = "Bajo peso"
	} else if IMC <= 24.9 {
		categoria = "Peso normal"
	} else if IMC <= 29.9 {
		categoria = "Sobrepeso"
	} else {
		categoria = "Obesidad"
	}

	fmt.Printf("La persona %s, de %d años de edad, tiene peso de %.2fkg y altura %.2fm. Su IMC es de %.2f, categorizado en %s\n", persona.Nombre, persona.Edad, persona.Peso, persona.Altura, IMC, categoria)
}

func IngresarPersona(nombre string, edad int, altura float64, peso float64, personas *[]Persona) *[]Persona {

	fmt.Println("Ingrese el nombre de usuario: ")
	fmt.Scan(&nombre)
	fmt.Println("Ingrese la edad: ")
	fmt.Scan(&edad)
	fmt.Println("Ingrese la altura: ")
	fmt.Scan(&altura)
	fmt.Println("Ingrese el peso: ")
	fmt.Scan(&peso)

	persona := CrearPersona(&nombre, &edad, &altura, &peso)
	*personas = append(*personas, *persona)

	fmt.Printf("===================== \nPersona creada exitosamente: \n%v\n=====================\n", persona)
	return personas
}
