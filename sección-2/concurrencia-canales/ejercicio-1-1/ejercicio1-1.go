/*
Ejercicio 1: Suma de números en un rango utilizando gorutinas
y canales Escribe un programa en Go que sume todos los números
en un rango dado (por ejemplo, de 1 a 100) utilizando gorutinas
y canales para dividir el trabajo entre varias gorutinas.
*/

package main

import (
	"fmt"
	"runtime"
	"sync"
)

func suma(first, last int, c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	s := 0
	for i := first; i <= last; i++ {
		s += i
	}
	c <- s
	fmt.Printf("La suma del %d hasta el %d es de: %d \n", first, last, s)
}
func suma2(first, last int) int {
	s := 0
	for i := first; i <= last; i++ {
		s += i
	}
	fmt.Printf("La suma del %d hasta el %d es de: %d \n", first, last, s)
	return s
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	var first, last, wgroups int
	var wg sync.WaitGroup
	var c = make(chan int)

	first = 10
	last = 100
	wgroups = 10

	groups := ((last - first + 1) / wgroups)

	for i := 0; i < wgroups; i++ {
		wg.Add(1)
		// go suma(groups*i+1, groups*(i+1), c, &wg)
		// go suma(first+i*groups, first+(i+1)*groups, c, &wg)
		// go suma(first+i*(groups+1), first+(i+1)*groups, c, &wg)
		go suma(first*(i+1), first+groups, c, &wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	totFinal := 0
	for v := range c {
		totFinal += v
	}
	fmt.Printf("La suma del %d hasta el %d es de: %d \n", first, last, totFinal)

	prueba := 0

	for i := first; i <= last; i++ {
		prueba += i
	}
	fmt.Println("La suma deberia dar:", prueba)

	// sumando := 0
	// for i := 0; i < groups; i++ {
	// 	suma2(groups*i+1, groups*(i+1))
	// 	sumando += suma2(groups*i+1, groups*(i+1))
	// }
	// fmt.Println("La f suma da:", sumando)

	// fmt.Println("La cantidad de grupos que tenemos son: ", groups)

}
