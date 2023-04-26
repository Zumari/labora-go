package main

import (
	"fmt"
	"sync"
)

func multiplicarMatrices(matrizA, matrizB [][]int, c chan [][]int, wg *sync.WaitGroup) {

	defer wg.Done()
	// Verificar si la multiplicación es posible
	if len(matrizA[0]) != len(matrizB) {
		panic("Las matrices no tienen el número adecuado de columnas y filas para multiplicarse")
	}

	// Definir la matriz resultado
	resultado := make([][]int, len(matrizA))
	for i := range resultado {
		resultado[i] = make([]int, len(matrizB[0]))
	}

	// Multiplicar matrices
	for i := 0; i < len(matrizA); i++ {
		for j := 0; j < len(matrizB[0]); j++ {
			for k := 0; k < len(matrizA[0]); k++ {
				resultado[i][j] += matrizA[i][k] * matrizB[k][j]
			}
		}
	}

	c <- resultado

}

func main() {
	var wg sync.WaitGroup
	var c = make(chan [][]int)

	var matrizA [][]int = [][]int{{1, 2, 3}, {3, 2, 1}}
	var matrizB [][]int = [][]int{{0, -1, -2}, {-2, 0, -1}, {-1, -2, 0}}

	wg.Add(1)
	go multiplicarMatrices(matrizA, matrizB, c, &wg)
	result := <-c
	fmt.Println(result)
}
