package main

import "fmt"

func main() {

	var a, b int
	fmt.Print("Ingrese el primer número: ")
	fmt.Scan(&a)
	fmt.Print("Ingrese el segundo número: ")
	fmt.Scan(&b)

	calcular(a, b)
}

func calcular(a int, b int) {

	var suma = a + b
	var resta = a - b
	var mult = a * b
	var div = float32(a) / float32(b)

	fmt.Printf("Suma: %v\n", suma)
	fmt.Printf("Resta: %v\n", resta)
	fmt.Printf("Multiplicación: %v\n", mult)
	fmt.Printf("División: %.2f", div)

}
