package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Fatalf("Error obteniendo el hostname: %v", err)
	}
	log.Printf("Listening on :%s\n", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "This is the pod %s\n", hostname)
		if err != nil {
			log.Printf("Error escribiendo la respuesta: %v", err)
		}
	})

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
