package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type ItemDetails struct {
	Item
	Details string `json:"details"`
}

var Items []Item

func handleRootResource(response http.ResponseWriter, request *http.Request) {
	msg := "Soy el recurso raiz"
	response.WriteHeader(http.StatusOK)
	response.Write([]byte(msg))
}

func handleOtherResource(response http.ResponseWriter, request *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte("respuesta fallida"))
		}
	}()
	msg := "Soy otro recurso"
	var item *Item = nil
	fmt.Println(item.ID) // KA BOOM
	response.Write([]byte(msg))
	response.WriteHeader(http.StatusOK)
}

func chatgptHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Printf("Endpointer, request from ip %s\n", request.RemoteAddr)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	inputStr := scanner.Text()
	//fmt.Fprintf(response, text)
	response.WriteHeader(http.StatusOK)
	_, err := response.Write([]byte(inputStr))
	if err != nil {
		fmt.Printf("error while writting bytes to response writer: %+v", err)
	}
}

func getItems(w http.ResponseWriter, r *http.Request) {
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")
	pageUser := r.URL.Query().Get("page")
	itemsUser := r.URL.Query().Get("itemsPerPage")
	// Convertir los parámetros a enteros
	page, err := strconv.Atoi(pageUser)
	if err != nil {
		page = 1
	}
	itemsPerPage, err := strconv.Atoi(itemsUser)
	if err != nil {
		itemsPerPage = 10
	}
	inicio := (page - 1) * itemsPerPage
	// Obtener los elementos del slice que corresponden a la página solicitada
	var resultado []Item
	if inicio >= 0 && inicio < len(Items) {
		final := inicio + itemsPerPage
		if final > len(Items) {
			final = len(Items)
		}
		resultado = Items[inicio:final]
	}
	// Función para obtener todos los elementos
	json.NewEncoder(w).Encode(resultado)
}

func getItemID(w http.ResponseWriter, r *http.Request) {

	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Obtenemos los parámetros de la URL utilizando Gorilla Mux
	params := mux.Vars(r)

	// Iteramos sobre cada elemento en 'Items' para buscar el que tenga el mismo ID que el proporcionado
	for _, item := range Items {
		if item.ID == params["id"] {
			// Si encontramos el elemento con el ID especificado, lo convertimos a JSON y lo enviamos como respuesta HTTP
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// Si no encontramos ningún elemento con el ID especificado, enviamos un objeto 'Item' vacío como respuesta HTTP
	json.NewEncoder(w).Encode(&Item{})

	// var idParam string = mux.Vars(r)["id"]
	// id, err := strconv.Atoi(idParam)
	// if err != nil {
	// 	w.WriteHeader(400)
	// 	w.Write([]byte("ID no puede ser convertido"))
	// 	return
	// }

	// if id >= len(Items)+1 {
	// 	w.WriteHeader(404)
	// 	w.Write([]byte("No se encuentra ningun item con ese ID"))
	// }

	// detalle := Items[id-1]
	// w.Header().Set("Contenido-Tipo", "aplicacion/json")
	// json.NewEncoder(w).Encode(detalle)

}

func getItemName(w http.ResponseWriter, r *http.Request) {
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")

	// Obtenemos los parámetros de la URL utilizando Gorilla Mux
	params := mux.Vars(r)

	// Iteramos sobre cada elemento en 'Items' para buscar el que tenga el mismo ID que el proporcionado

	var wanted []Item
	count := 0
	for _, item := range Items {
		if strings.EqualFold(item.Name, params["name"]) {
			// Si encontramos el elemento con el ID especificado, lo convertimos a JSON y lo enviamos como respuesta HTTP
			wanted = append(wanted, item)
			count++
		}
	}

	if count > 1 {
		json.NewEncoder(w).Encode(wanted)
		return
	}

	// Si no encontramos ningún elemento con el ID especificado, enviamos un objeto 'Item' vacío como respuesta HTTP
	json.NewEncoder(w).Encode(&Item{})

}

func createItem(w http.ResponseWriter, r *http.Request) {
	// Recibimos la información
	var newItem Item
	reqBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Fprintf(w, "Ingrese un Item válido")
	}

	// Asignamos la informacion a la variable newItem
	json.Unmarshal(reqBody, &newItem)

	// newItem.ID = string(len(Items) + 1)
	newItem.ID = fmt.Sprintf("item%d", len(Items)+1)
	fmt.Println(string(len(Items) + 1))
	Items = append(Items, newItem)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// Función para actualizar un elemento existente
	var UpdateItem Item
	parametros := mux.Vars(r)
	rqBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Inserte un item válido")
		return
	}
	json.Unmarshal(rqBody, &UpdateItem)
	for i, item := range Items {
		if item.ID == parametros["id"] {
			Items = append(Items[:i], Items[i+1:]...)
			UpdateItem.ID = parametros["id"]
			Items = append(Items, UpdateItem)
		}
	}
	// Enviar una respuesta exitosa al cliente
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(UpdateItem)
	fmt.Println("¡Item actualizado!")
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Establecemos el encabezado "Content-Type" de la respuesta HTTP como "application/json"
	w.Header().Set("Content-Type", "application/json")
	// Función para eliminar un elemento
	parametros := mux.Vars(r)
	for i, item := range Items {
		if item.ID == parametros["id"] {
			Items = append(Items[:i], Items[i+1:]...)
			fmt.Fprintf(w, "El item con el ID %s fue eliminada", item.ID)
			return
		}
	}
	fmt.Fprintf(w, "Usuario no encontrado")

}

func getItemDetails(id string) ItemDetails {
	// Simula la obtención de detalles desde una fuente externa con un time.Sleep
	time.Sleep(100 * time.Millisecond)
	var foundItem Item
	for _, item := range Items {
		if item.ID == string(id) {
			foundItem = item
			break
		}
	}
	//Obviamente, aquí iria un SELECT si es SQL o un llamado a un servicio externo
	//pero esta busqueda del item junto con Details, la hacemos a mano.
	return ItemDetails{
		Item:    foundItem,
		Details: fmt.Sprintf("Detalles para el item %d", id),
	}
}

func main() {

	// samuel := Item{"01", "samuel"}
	// pedro := Item{"02", "pedro"}
	// pablo := Item{"03", "pablo"}
	// maria := Item{"04", "maria"}
	// sol := Item{"05", "sol"}
	// eunice := Item{"06", "eunice"}
	// camila := Item{"07", "camila"}
	// bryan := Item{"08", "samuel"}
	// lester := Item{"09", "lester"}
	ignacio := Item{"item11", "Item 1"}
	anibal := Item{"item12", "Item 1"}
	Items = append(Items, ignacio, anibal)

	// Items = []Item{samuel, pedro, pablo, maria, sol, eunice, camila, bryan, lester, ignacio}

	for i := 1; i <= 10; i++ {
		Items = append(Items, Item{ID: fmt.Sprintf("item%d", i), Name: fmt.Sprintf("Item %d", i)})
	}

	router := mux.NewRouter() // aca creamos el coso para configurar el servidor, el coso se llama router

	// aca definimos el comportamiento del servidor
	router.HandleFunc("/", handleRootResource).Methods("GET")
	router.HandleFunc("/otro", handleOtherResource).Methods("GET")
	router.HandleFunc("/chatgpt", chatgptHandler).Methods("GET")

	//Tarea 25/04/2023
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/id/{id}", getItemID).Methods("GET")

	// Nuevos endpoints 26/04/2023
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")
	router.HandleFunc("/items/name/{name}", getItemName).Methods("GET")

	// aca termine de definir el comportamiento

	// levantar el servidor en un "puerto"
	var portNumber int = 9999
	fmt.Println("Listen in port ", portNumber)
	err := http.ListenAndServe(":"+strconv.Itoa(portNumber), router)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ADIOS")
}
