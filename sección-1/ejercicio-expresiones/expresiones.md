# Enuciado

Expresar a un número cualquiera como la suma de una serie de números siguiendo las restricciones impuestas a continuación.

Sea x el número.
x = s1 + s2 + s3 + s4 + s5

Donde
0 < s1 < 50
0 < s2 < 50
0 < s3 < 600
0 < s4 < 800
0 < s5 < (Infinito)

De esta forma, si X vale 120, entonces
s1 = 50, s2 = 50, s3 = 20, s4 = 0 y s5 = 0

Si X vale 860, entonces
s1 = 50, s2 = 50, s3 = 600, s4 = 160 y s5 = 0

* Este ejercicio lo tomé de una guia de mi facultad y el enunciado terminaba acá recuerdo que el enunciado no me alcanzaba para entender bien como encarar esto de modo que estimo conveniente agragegar un poco más de contexto para resolver esto *

**=== CONTEXTO ADICIONAL ====**

Lo que pueden hacer es crearse un solo archivo que contenga la solución en código


Por ejemplo:
package main

import (
	"fmt"
)

func main() {
	var x int = 120 // aca hacemos que x valga lo que queremos que valga
	segmentarValorPorRangos(x)
}

func segmentarValorPorRangos(x int) {
	var s1,s2,s3,s4,s5 int
	// acá implementamos la magia de ir segmentando ese valor, que va a llegar como argumento a esta funcion
  // tip: (no lo miren si quieren realmetne PENSAR por su cuenta sola......................................... el tip es medio pedorro igual, el tip es hay que usar la sentencia condicional IF varias veces :D )

	fmt.Println(s1,s2,s3,s4,s5) // por ahora pueden imprimir en pantalla... más adelante veremos que diferencia hay entre hacer esto y devovler como retorno de la función a estas 5 variables!
}

