package main

import (
	"fmt"
	"sync"
	"time"
)

func suma(n []int, total *int, wg *sync.WaitGroup) {

	defer wg.Done()

	*total = 0
	for _, suma := range n {
		*total += suma
	}

}

func sumaSecuencial(n []int, total *int) {

	*total = 0
	for _, suma := range n {
		*total += suma
	}

}

func main() {

	var suma1, suma2, sumaSec int
	var wg sync.WaitGroup

	wg.Add(2)

	numeros := make([]int, 100000000000)
	for i := 0; i < len(numeros); i++ {
		numeros[i] = i
	}

	//Sumando de manera concurrente

	inicio2 := time.Now()
	mitad := len(numeros) / 2
	go suma(numeros[:mitad], &suma1, &wg)
	go suma(numeros[mitad:], &suma2, &wg)

	wg.Wait()
	sumaTotal := suma1 + suma2
	fmt.Println("La suma total es: ", sumaTotal)
	duracion := time.Since(inicio2)
	fmt.Println("La duracion de manera concurrente es: ", duracion)

	inicio1 := time.Now()
	sumaSecuencial(numeros, &sumaSec)

	wg.Wait()
	fmt.Println("La suma total es: ", sumaTotal)
	duracion1 := time.Since(inicio1)
	fmt.Println("La duracion de manera secuencial es: ", duracion1)

}
