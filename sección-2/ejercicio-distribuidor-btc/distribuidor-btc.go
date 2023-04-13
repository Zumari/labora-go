/*
Tienes 50 bitcoins para distribuir a 10 usuarios: Matthew, Sarah,
Augustus, Heidi, Emilie, Peter, Giana, Adriano, Aaron, Elizabeth
Las monedas se distribuirán en función de las vocales contenidas en cada nombre donde:

a: 1 moneda e: 1 moneda i: 2 monedas o: 3 monedas u: 4 monedas

y un usuario no puede obtener más de 10 monedas.Imprime un mapa
con el nombre de cada usuario y la cantidad de monedas repartidas.
Después de distribuir todas las monedas, te deben quedar 2 monedas.
*/
package main

import (
	"fmt"
)

func distribuir(users []string, coins *int, distribution map[string]int) *map[string]int {
	var user string
	var sum int
	vocales := map[string]int{
		"a": 1,
		"e": 1,
		"i": 2,
		"o": 3,
		"u": 4,
		"A": 1,
		"E": 1,
		"I": 2,
		"O": 3,
		"U": 4,
	}

	for _, name := range users {
		user = name
		sum = 0
		for _, letter := range user {
			lt := string(letter)
			if value, ok := vocales[lt]; ok {
				fmt.Printf("letra encontrada %s en %s con valor de %d\n", lt, user, value)
				sum += value
				fmt.Printf("la suma va en %d\n", sum)
			}
		}
		if sum < 10 {
			distribution[name] = sum
			*coins -= sum
		} else {
			distribution[name] = 10
			*coins -= 10
		}

	}

	return &distribution
}

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	distribuir(users, &coins, distribution)
	fmt.Println(distribution)
	fmt.Println("Coins left:", coins)
}
