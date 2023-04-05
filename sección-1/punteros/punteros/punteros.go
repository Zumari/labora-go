package main

import "fmt"

func main() {

	var a, b = 10, 20
	var ptrA *int = &a

	fmt.Printf("Valores iniciales: a = %v, b= %v\n", a, b)

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("*ptrA = ", *ptrA)
	fmt.Println("*ptrA = ", &ptrA)

	a = *ptrA

	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	fmt.Println("*ptrA = ", *ptrA)
	fmt.Println("*ptrA = ", &ptrA)
	fmt.Printf("Valores finales: a= %v, b= %v", *ptrA, b)

}
