package main

import "fmt"

func main() {

	var a, b = 10, 20
	var ptrA *int = &a

	fmt.Printf("Valores iniciales: a = %v, b= %v\n", a, b)

	a = *ptrA + b
	b = *ptrA - b
	a = *ptrA - b

	fmt.Printf("Valores finales: a= %v, b= %v", a, b)

}
