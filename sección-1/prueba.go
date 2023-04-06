package main

import (
	"fmt"
)

func main() {
	// matriz := [][]int{{1, 2}, {3, 4}, {5, 6}}
	// for i := 0; i < len(matriz); i++ {
	// 	for j := 0; j < len(matriz[i]); j++ {
	// 		fmt.Printf("%d ", matriz[i][j])
	// 	}
	// 	fmt.Println()
	// }

	numeros := []int{1, 2, 3, 4, 5}

	for _, numero := range numeros {
		fmt.Println(numero)
	}

}
