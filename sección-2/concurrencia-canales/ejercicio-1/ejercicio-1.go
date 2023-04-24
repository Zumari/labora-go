/*
Ejercicio 1: Suma de números en un rango utilizando gorutinas
y canales Escribe un programa en Go que sume todos los números
en un rango dado (por ejemplo, de 1 a 100) utilizando gorutinas
y canales para dividir el trabajo entre varias gorutinas.
*/

package main

import (
	"fmt"
	"sync"
)

func suma(n []int, c chan int, wg *sync.WaitGroup) {

	defer wg.Done()

	s := 0

	for _, v := range n {
		s += v
	}

	c <- s
}

func main() {

	var sum1, sum2 int
	var numeros = make([]int, 100)
	var channel = make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < len(numeros); i++ {
		numeros[i] = i + 1
	}

	wg.Add(2)

	mitadNumeros := len(numeros) / 2
	go suma(numeros[:mitadNumeros], channel, &wg)
	go suma(numeros[mitadNumeros:], channel, &wg)

	sum1, sum2 = <-channel, <-channel
	go func() {
		wg.Wait()
		close(channel)
	}()
	sumaTotal := sum1 + sum2

	fmt.Println("La suma1 de numeros es: ", sum1)
	fmt.Println("La suma2 de numeros es: ", sum2)
	fmt.Println("La suma de numeros es: ", sumaTotal)

}
