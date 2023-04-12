package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	nombre string
	edad   int
	altura float64
	peso   float64
}

func crearPersona(nombre *string, edad *int, altura *float64, peso *float64) *Persona {

	if *edad < 0 || *altura < 0 || *peso < 0 {
		fmt.Println("Error en alguno de los datos ingresados")
		return nil
	}
	return &Persona{nombre: *nombre, edad: *edad, altura: *altura, peso: *peso}

}

func ordenarPersonas(personas []Persona, criterio int) []Persona {
	switch criterio {
	case 1:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].nombre < personas[j].nombre
		})
	case 2:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].edad < personas[j].edad
		})
	case 3:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].altura < personas[j].altura
		})
	case 4:
		sort.Slice(personas, func(i, j int) bool {
			return personas[i].peso < personas[j].peso
		})
	default:
	}

	return personas
}

func buscarPersona(personas []Persona, criterio int, valor any) {
	var comparar func(p Persona) bool
	switch criterio {
	case 1:
		comparar = func(p Persona) bool {
			return p.nombre == valor.(string)
		}
	case 2:
		comparar = func(p Persona) bool {
			return p.edad == valor.(int)
		}
	case 3:
		comparar = func(p Persona) bool {
			return p.altura == valor.(float64)
		}
	case 4:
		comparar = func(p Persona) bool {
			return p.peso == valor.(float64)
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

func calcularIMC(persona Persona) {
	IMC := persona.peso / (persona.altura * persona.altura)
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

	fmt.Printf("La persona %s, de %d años de edad, tiene peso de %.2fkg y altura %.2fm. Su IMC es de %.2f, categorizado en %s\n", persona.nombre, persona.edad, persona.peso, persona.altura, IMC, categoria)
}

func ingresarPersona(nombre string, edad int, altura float64, peso float64, personas *[]Persona) *[]Persona {

	fmt.Println("Ingrese el nombre de usuario: ")
	fmt.Scan(&nombre)
	fmt.Println("Ingrese la edad: ")
	fmt.Scan(&edad)
	fmt.Println("Ingrese la altura: ")
	fmt.Scan(&altura)
	fmt.Println("Ingrese el peso: ")
	fmt.Scan(&peso)

	persona := crearPersona(&nombre, &edad, &altura, &peso)
	*personas = append(*personas, *persona)

	fmt.Printf("===================== \nPersona creada exitosamente: \n%v\n=====================\n", persona)
	return personas
}

func main() {

	var nombre string
	var edad int
	var altura float64
	var peso float64
	var personas []Persona
	var accion, criterio int
	i := 0
	for i != 1 {

		fmt.Printf("Ingrese el proceso que desea hacer: \n1. Agregar persona\n2. Ordenar personas \n3. Buscar persona \n4. Calcular imc \n5. Cerrar\n")

		fmt.Scan(&accion)

		switch accion {
		case 1:
			ingresarPersona(nombre, edad, altura, peso, &personas)
			fmt.Printf("\nGrupo de personas: \n%v\n", personas)
		case 2: // Ordenar personas
			fmt.Println("Criterios desea ordenar a las personas: \n1. Nombre \n2. Edad \n3. Altura \n4. Peso")
			fmt.Println("Escriba el número de criterio: ")
			fmt.Scan(&criterio)
			fmt.Printf("\nPersonas ordenadas: \n%v\n", ordenarPersonas(personas, criterio))

		case 3: // Buscar personas
			fmt.Println("¿Mediante cúal criterio desea buscar a las personas? \n1. Nombre \n2. Edad \n3. Altura \n4. Peso")
			fmt.Scan(&criterio)
			fmt.Println("Ingrese el atributo a buscar: ")
			switch criterio {
			case 1:
				var valor string
				fmt.Scan(&valor)
				buscarPersona(personas, criterio, valor)
			case 2:
				var valor int
				fmt.Scan(&valor)
				buscarPersona(personas, criterio, valor)
			case 3:
				var valor float64
				fmt.Scan(&valor)
				buscarPersona(personas, criterio, valor)
			case 4:
				var valor float64
				fmt.Scan(&valor)
				buscarPersona(personas, criterio, valor)
			}
		case 4:
			var personaIMC int
			fmt.Printf("Ingrese el indice de la persona a la que le desea calcular el IMC \n%v\n", personas)
			fmt.Scan(&personaIMC)
			if personaIMC > len(personas) {
				fmt.Println("Indice inválido,intentelo de nuevo")
			} else {
				calcularIMC(personas[personaIMC])
			}
		case 5:
			i = 1
		default:
			fmt.Println("Opción inválida")
		}
	}
}
