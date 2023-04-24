package main

import (
	"fmt"
	"sync"
	"time"
)

// Define un tipo de mensaje que tenga un tipo y un contenido
type Message struct {
	Type    string
	Content string
}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	// Define los canales y las gorutinas para cada tipo de mensaje
	emailCh := make(chan Message)
	go processMessages("email", emailCh, &wg)
	smsCh := make(chan Message)
	go processMessages("SMS", smsCh, &wg)
	pushCh := make(chan Message)
	go processMessages("notificación push", pushCh, &wg)
	// Agrega algunos mensajes a la cola
	emailCh <- Message{"email", "Este es un correo electrónico de prueba."}
	smsCh <- Message{"SMS", "Este es un mensaje de texto de prueba."}
	pushCh <- Message{"notificación push", "Esta es una notificación push de prueba."}
	wg.Wait()
}

// Procesa los mensajes en orden de llegada para un tipo específico de mensaje
func processMessages(msgType string, msgCh chan Message, wg *sync.WaitGroup) {
	msg := <-msgCh
	fmt.Printf("Procesando mensaje %q de tipo %q...\n", msg.Content, msg.Type)
	switch msgType {
	case "email":
		time.Sleep(2 * time.Second)
		fmt.Printf("El mensaje %q de tipo %q ya se procesó\n", msg.Content, msg.Type)
	case "SMS":
		time.Sleep(3 * time.Second)
		fmt.Printf("El mensaje %q de tipo %q ya se procesó\n", msg.Content, msg.Type)
	case "notificación push":
		time.Sleep(1 * time.Second)
		fmt.Printf("El mensaje %q de tipo %q ya se procesó\n", msg.Content, msg.Type)
	}
	defer wg.Done()
}

/*

func main() {
	emailChan := make(chan Message)
	signalChannel := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		var seguirProcesando bool = true
		for seguirProcesando {
			select {
			case msg, ok := <-emailChan:
				if !ok {
					fmt.Println("Canal cerrado")
					seguirProcesando = false
					break
				}
				fmt.Printf("Procesando mensaje de notificación correo electronico: %s\n", msg.Content)
				time.Sleep(3 * time.Second)
				fmt.Printf("se termino de procesar el mensaje correo electronico: %s\n", msg.Content)
				fmt.Println("Envio señal de continuar")
				signalChannel <- true
			}
		}
	}()
	go func(emailChan chan<- Message) {
		scanner := bufio.NewScanner(os.Stdin)
		for true {
			// paremos acá!
			fmt.Print("Ingrese su mensaje o presione CTR+D para salir:")
			hasMoreInput := scanner.Scan()
			if !hasMoreInput {
				fmt.Println("SALIENDO")
				break
			}
			inputStr := scanner.Text()
			message := Message{Type: "email", Content: inputStr}
			emailChan <- message

			var signal bool
			signal = <-signalChannel
			fmt.Println("Llego señal de continuar", signal)
		}
		close(emailChan)
	}(emailChan)
	wg.Wait()
	fmt.Println("Todos los mensajes han sido procesados.")
}

*/
