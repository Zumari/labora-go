package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// handle route using handler function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to new server, Mari!")
	})

	// listen to port
	http.ListenAndServe(":5050", nil)

	// Crear el enrutador y  luego definir las rutas en la función definirRutas

	enrutador := mux.NewRouter()

	enrutador.HandleFunc("/hola", func(respuesta http.ResponseWriter, peticion *http.Request) {
		respuesta.Write([]byte("¡Hola, mundo!"))
		respuesta.WriteHeader(http.StatusOK)
	}).Methods("GET")

}
